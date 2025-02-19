package ws

import "net/http"

// WebSocketHandler 定义WebSocket处理器接口
type WebSocketHandler interface {
	// HandleConnection 处理新的WebSocket连接
	HandleConnection(conn Connection)
}

// Connection 定义WebSocket连接接口
type Connection interface {
	// GetID 获取连接唯一标识
	GetID() string
	// Send 发送消息
	Send(messageType int, data []byte) error
	// Close 关闭连接
	Close() error
	// GetRequest 获取原始HTTP请求
	GetRequest() *http.Request
}

// MessageHandler 定义消息处理器接口
type MessageHandler interface {
	// OnMessage 处理接收到的消息
	OnMessage(conn Connection, messageType int, data []byte)
	// OnError 处理错误
	OnError(conn Connection, err error)
	// OnClose 处理连接关闭
	OnClose(conn Connection)
}

// ConnectionManager 定义连接管理器接口
type ConnectionManager interface {
	// AddConnection 添加新连接
	AddConnection(conn Connection)
	// RemoveConnection 移除连接
	RemoveConnection(conn Connection)
	// GetConnection 获取指定连接
	GetConnection(id string) Connection
	// Broadcast 广播消息给所有连接
	Broadcast(messageType int, data []byte)
}