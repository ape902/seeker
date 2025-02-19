package worker_pool

// Config 告警引擎配置
type Config struct {
	// 数据源配置
	DataSourceConfig struct {
		// 数据源连接超时时间（秒）
		ConnectionTimeout int `json:"connection_timeout" yaml:"connection_timeout"`
		// 数据源请求超时时间（秒）
		RequestTimeout int `json:"request_timeout" yaml:"request_timeout"`
		// 数据源健康检查间隔（秒）
		HealthCheckInterval int `json:"health_check_interval" yaml:"health_check_interval"`
	} `json:"data_source" yaml:"data_source"`

	// 规则管理配置
	RuleConfig struct {
		// 规则同步间隔（秒）
		SyncInterval int `json:"sync_interval" yaml:"sync_interval"`
		// 规则评估超时时间（秒）
		EvaluationTimeout int `json:"evaluation_timeout" yaml:"evaluation_timeout"`
	} `json:"rule" yaml:"rule"`

	// 事件管理配置
	EventConfig struct {
		// 事件队列大小
		QueueSize int `json:"queue_size" yaml:"queue_size"`
		// 事件处理超时时间（秒）
		ProcessTimeout int `json:"process_timeout" yaml:"process_timeout"`
		// 事件保留时间（天）
		RetentionDays int `json:"retention_days" yaml:"retention_days"`
	} `json:"event" yaml:"event"`

	// 工作池配置
	WorkerPoolConfig struct {
		// 工作池大小
		Size int `json:"size" yaml:"size"`
		// 任务队列大小
		QueueSize int `json:"queue_size" yaml:"queue_size"`
	} `json:"worker_pool" yaml:"worker_pool"`
}

// NewDefaultConfig 创建默认配置
func NewDefaultConfig() *Config {
	config := &Config{}

	// 设置数据源默认配置
	config.DataSourceConfig.ConnectionTimeout = 10
	config.DataSourceConfig.RequestTimeout = 30
	config.DataSourceConfig.HealthCheckInterval = 60

	// 设置规则管理默认配置
	config.RuleConfig.SyncInterval = 300
	config.RuleConfig.EvaluationTimeout = 30

	// 设置事件管理默认配置
	config.EventConfig.QueueSize = 1000
	config.EventConfig.ProcessTimeout = 30
	config.EventConfig.RetentionDays = 30

	// 设置工作池默认配置
	config.WorkerPoolConfig.Size = 10
	config.WorkerPoolConfig.QueueSize = 1000

	return config
}
