package worker_pool

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ape902/corex/logx"
)

// DataSourceType 数据源类型
type DataSourceType string

const (
	PrometheusDataSource    DataSourceType = "prometheus"
	ElasticsearchDataSource DataSourceType = "elasticsearch"
)

// DataSourceManager 数据源管理器接口
type DataSourceManager interface {
	Start(ctx context.Context) error
	Stop()
	GetData(sourceID string) (interface{}, error)
	RegisterDataSource(source DataSource) error
}

// DataSource 数据源接口
type DataSource interface {
	GetType() DataSourceType
	GetID() string
	Query(query string) (interface{}, error)
}

// dataSourceManagerImpl 数据源管理器实现
type dataSourceManagerImpl struct {
	sources map[string]DataSource
	config  *Config
	mux     sync.RWMutex
	quit    chan struct{}
}

// NewDataSourceManager 创建数据源管理器
func NewDataSourceManager(config *Config) DataSourceManager {
	return &dataSourceManagerImpl{
		sources: make(map[string]DataSource),
		config:  config,
		quit:    make(chan struct{}),
	}
}

// Start 启动数据源管理器
func (m *dataSourceManagerImpl) Start(ctx context.Context) error {
	logx.Info("Starting data source manager...")
	// 启动健康检查
	go m.healthCheck(ctx)

	return nil
}

// Stop 停止数据源管理器
func (m *dataSourceManagerImpl) Stop() {
	logx.Info("Stopping data source manager...")
	close(m.quit)
}

// GetData 从指定数据源获取数据
func (m *dataSourceManagerImpl) GetData(sourceID string) (interface{}, error) {
	m.mux.RLock()
	defer m.mux.RUnlock()

	source, exists := m.sources[sourceID]
	if !exists {
		return nil, ErrDataSourceNotFound
	}

	// 尝试查询数据，如果失败则返回连接错误
	data, err := source.Query("")
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDataSourceConnectionFailed, err)
	}

	return data, nil
}

// RegisterDataSource 注册数据源
func (m *dataSourceManagerImpl) RegisterDataSource(source DataSource) error {
	if source == nil {
		return ErrInvalidDataSourceType
	}

	m.mux.Lock()
	defer m.mux.Unlock()

	if _, exists := m.sources[source.GetID()]; exists {
		return ErrDataSourceAlreadyExists
	}

	m.sources[source.GetID()] = source
	logx.Info("Registered data source: %s", source.GetID())
	return nil
}

// healthCheck 数据源健康检查
func (m *dataSourceManagerImpl) healthCheck(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-m.quit:
			return
		case <-ticker.C:
			m.checkDataSources()
		}
	}
}

// checkDataSources 检查所有数据源的健康状态
func (m *dataSourceManagerImpl) checkDataSources() {
	m.mux.RLock()
	defer m.mux.RUnlock()

	for id, source := range m.sources {
		if _, err := source.Query(""); err != nil {
			logx.Errorf("Data source %s health check failed: %v", id, err)
		}
	}
}
