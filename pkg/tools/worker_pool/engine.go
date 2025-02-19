package worker_pool

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/ape902/corex/logx"
)

// AlertEngine 告警引擎核心结构体
type AlertEngine struct {
	dataSourceManager DataSourceManager
	ruleManager       RuleManager
	eventManager      EventManager
	workerPool        *WorkerPool
	quit              chan struct{}
	wg                sync.WaitGroup
}

// NewAlertEngine 创建新的告警引擎实例
func NewAlertEngine(config *Config) *AlertEngine {
	if config == nil {
		config = NewDefaultConfig()
	}

	return &AlertEngine{
		dataSourceManager: NewDataSourceManager(config),
		ruleManager:       NewRuleManager(config),
		eventManager:      NewEventManager(config),
		workerPool:        NewWorkerPool(config.WorkerPoolConfig.Size),
		quit:              make(chan struct{}),
	}
}

// Start 启动告警引擎
func (e *AlertEngine) Start(ctx context.Context) error {
	logx.Info("Starting alert engine...")

	// 创建一个错误通道用于收集启动过程中的错误
	errChan := make(chan error, 1)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// 启动数据源管理器
	if err := e.dataSourceManager.Start(ctx); err != nil {
		logx.Errorf("Failed to start data source manager: %v", err)
		return fmt.Errorf("failed to start data source manager: %v", err)
	}

	// 启动规则管理器
	if err := e.ruleManager.Start(ctx); err != nil {
		logx.Error("Failed to start rule manager: %v", err)
		e.dataSourceManager.Stop() // 回滚已启动的组件
		return fmt.Errorf("failed to start rule manager: %v", err)
	}

	// 启动事件管理器
	if err := e.eventManager.Start(ctx); err != nil {
		logx.Error("Failed to start event manager: %v", err)
		e.ruleManager.Stop() // 回滚已启动的组件
		e.dataSourceManager.Stop()
		return fmt.Errorf("failed to start event manager: %v", err)
	}

	// 启动工作池
	e.workerPool.Start()

	// 启动主处理循环
	e.wg.Add(1)
	go func() {
		defer e.wg.Done()
		if err := e.run(ctx); err != nil {
			errChan <- err
		}
	}()

	// 检查是否有启动错误
	select {
	case err := <-errChan:
		e.Stop()
		return fmt.Errorf("alert engine failed to start: %v", err)
	case <-time.After(time.Second): // 给予一定的启动时间
		logx.Info("Alert engine started successfully")
		return nil
	}
}

// Stop 停止告警引擎
func (e *AlertEngine) Stop() {
	logx.Info("Stopping alert engine...")
	close(e.quit)

	// 等待所有goroutine结束
	e.wg.Wait()

	// 按照依赖关系的反序停止各个组件
	e.workerPool.Stop()
	e.eventManager.Stop()
	e.ruleManager.Stop()
	e.dataSourceManager.Stop()

	logx.Info("Alert engine stopped successfully")
}

// run 主处理循环
func (e *AlertEngine) run(ctx context.Context) error {
	ticker := time.NewTicker(time.Second * 10)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-e.quit:
			return nil
		case <-ticker.C:
			if err := e.processRules(); err != nil {
				logx.Errorf("Failed to process rules: %v", err)
			}
		}
	}
}

// processRules 处理规则
func (e *AlertEngine) processRules() error {
	rules := e.ruleManager.GetActiveRules()
	if len(rules) == 0 {
		return nil
	}

	var wg sync.WaitGroup
	errors := make(chan error, len(rules))

	for _, rule := range rules {
		wg.Add(1)
		ruleCopy := rule // 创建规则的副本以避免闭包问题
		e.workerPool.Submit(func() {
			defer wg.Done()
			if err := e.evaluateRule(ruleCopy); err != nil {
				logx.Errorf("Failed to evaluate rule %s: %v", ruleCopy.GetID(), err)
				errors <- fmt.Errorf("rule %s evaluation failed: %v", ruleCopy.GetID(), err)
			}
		})
	}

	// 等待所有规则评估完成
	go func() {
		wg.Wait()
		close(errors)
	}()

	// 收集错误
	var errs []error
	for err := range errors {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return fmt.Errorf("multiple rule evaluation errors: %v", errs)
	}

	return nil
}

// evaluateRule 评估单个规则
func (e *AlertEngine) evaluateRule(rule Rule) error {
	// 从数据源获取数据
	data, err := e.dataSourceManager.GetData(rule.GetDataSourceID())
	if err != nil {
		return fmt.Errorf("failed to get data from source: %v", err)
	}

	// 评估规则条件
	matched, err := rule.Evaluate(data)
	if err != nil {
		return fmt.Errorf("failed to evaluate rule: %v", err)
	}

	// 如果规则匹配，创建告警事件
	if matched {
		event := rule.CreateEvent(data)
		if err := e.eventManager.HandleEvent(event); err != nil {
			return fmt.Errorf("failed to handle event: %v", err)
		}
	}

	return nil
}
