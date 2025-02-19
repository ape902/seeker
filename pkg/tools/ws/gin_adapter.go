package ws

import (
	"fmt"
	"net/http"
	"sync/atomic"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// GinWebSocketHandler Gin框架的WebSocket处理器
type GinWebSocketHandler struct {
	connManager    ConnectionManager
	messageHandler MessageHandler
	connCounter    uint64
}

// NewGinWebSocketHandler 创建新的Gin WebSocket处理器
func NewGinWebSocketHandler(messageHandler MessageHandler) *GinWebSocketHandler {
	return &GinWebSocketHandler{
		connManager:    NewConnectionManager(),
		messageHandler: messageHandler,
	}
}

// Handle 处理Gin的WebSocket请求
func (h *GinWebSocketHandler) Handle(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("WebSocket upgrade error: %v", err))
		return
	}

	// 创建新的连接
	conn := &ginConnection{
		id:     fmt.Sprintf("conn_%d", atomic.AddUint64(&h.connCounter, 1)),
		conn:   ws,
		request: c.Request,
	}

	// 添加到连接管理器
	h.connManager.AddConnection(conn)

	// 启动消息处理
	go h.handleConnection(conn)
}

// handleConnection 处理WebSocket连接
func (h *GinWebSocketHandler) handleConnection(conn *ginConnection) {
	defer func() {
		h.connManager.RemoveConnection(conn)
		conn.Close()
		if h.messageHandler != nil {
			h.messageHandler.OnClose(conn)
		}
	}()

	for {
		messageType, message, err := conn.conn.ReadMessage()
		if err != nil {
			if h.messageHandler != nil {
				h.messageHandler.OnError(conn, err)
			}
			break
		}

		if h.messageHandler != nil {
			h.messageHandler.OnMessage(conn, messageType, message)
		}
	}
}

// ginConnection Gin框架的WebSocket连接实现
type ginConnection struct {
	id      string
	conn    *websocket.Conn
	request *http.Request
}

func (c *ginConnection) GetID() string {
	return c.id
}

func (c *ginConnection) Send(messageType int, data []byte) error {
	return c.conn.WriteMessage(messageType, data)
}

func (c *ginConnection) Close() error {
	return c.conn.Close()
}

func (c *ginConnection) GetRequest() *http.Request {
	return c.request
}