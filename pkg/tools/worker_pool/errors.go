package worker_pool

import "errors"

var (
	// ErrDataSourceNotFound 数据源未找到错误
	ErrDataSourceNotFound = errors.New("data source not found")

	// ErrDataSourceAlreadyExists 数据源已存在错误
	ErrDataSourceAlreadyExists = errors.New("data source already exists")

	// ErrRuleNotFound 规则未找到错误
	ErrRuleNotFound = errors.New("rule not found")

	// ErrRuleAlreadyExists 规则已存在错误
	ErrRuleAlreadyExists = errors.New("rule already exists")

	// ErrEventQueueFull 事件队列已满错误
	ErrEventQueueFull = errors.New("event queue is full")

	// ErrInvalidDataSourceType 无效的数据源类型错误
	ErrInvalidDataSourceType = errors.New("invalid data source type")

	// ErrInvalidRuleConfig 无效的规则配置错误
	ErrInvalidRuleConfig = errors.New("invalid rule configuration")

	// ErrDataSourceConnectionFailed 数据源连接失败错误
	ErrDataSourceConnectionFailed = errors.New("data source connection failed")

	// ErrRuleEvaluationFailed 规则评估失败错误
	ErrRuleEvaluationFailed = errors.New("rule evaluation failed")

	// ErrEventHandlingFailed 事件处理失败错误
	ErrEventHandlingFailed = errors.New("event handling failed")
)
