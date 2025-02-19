package ws

import (
	"sync"
)

// DefaultConnectionManager 默认的连接管理器实现
type DefaultConnectionManager struct {
	connections sync.Map
}

// NewConnectionManager 创建新的连接管理器
func NewConnectionManager() ConnectionManager {
	return &DefaultConnectionManager{}
}

// AddConnection 添加新连接
func (m *DefaultConnectionManager) AddConnection(conn Connection) {
	m.connections.Store(conn.GetID(), conn)
}

// RemoveConnection 移除连接
func (m *DefaultConnectionManager) RemoveConnection(conn Connection) {
	m.connections.Delete(conn.GetID())
}

// GetConnection 获取指定连接
func (m *DefaultConnectionManager) GetConnection(id string) Connection {
	if conn, ok := m.connections.Load(id); ok {
		return conn.(Connection)
	}
	return nil
}

// Broadcast 广播消息给所有连接
func (m *DefaultConnectionManager) Broadcast(messageType int, data []byte) {
	m.connections.Range(func(key, value interface{}) bool {
		if conn, ok := value.(Connection); ok {
			_ = conn.Send(messageType, data)
		}
		return true
	})
}