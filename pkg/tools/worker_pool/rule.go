package worker_pool

import (
	"context"
	"sync"
	"time"

	"github.com/ape902/corex/logx"
)

// Rule 告警规则接口
type Rule interface {
	GetID() string
	GetDataSourceID() string
	Evaluate(data interface{}) (bool, error)
	CreateEvent(data interface{}) Event
}

// RuleManager 规则管理器接口
type RuleManager interface {
	Start(ctx context.Context) error
	Stop()
	GetActiveRules() []Rule
	UpdateRules(rules []Rule) error
}

// ruleManagerImpl 规则管理器实现
type ruleManagerImpl struct {
	rules  []Rule
	config *Config
	mux    sync.RWMutex
	quit   chan struct{}
}

// NewRuleManager 创建规则管理器
func NewRuleManager(config *Config) RuleManager {
	return &ruleManagerImpl{
		rules:  make([]Rule, 0),
		config: config,
		quit:   make(chan struct{}),
	}
}

// Start 启动规则管理器
func (m *ruleManagerImpl) Start(ctx context.Context) error {
	logx.Info("Starting rule manager...")

	// 启动规则同步
	go m.syncRules(ctx)

	return nil
}

	// Stop 停止规则管理器
	func (m *ruleManagerImpl) Stop() {
		logx.Info("Stopping rule manager...")
		close(m.quit)
	}

	// GetActiveRules 获取所有活跃的规则
	func (m *ruleManagerImpl) GetActiveRules() []Rule {
		m.mux.RLock()
		defer m.mux.RUnlock()

		rules := make([]Rule, len(m.rules))
		copy(rules, m.rules)
		return rules
	}

	// UpdateRules 更新规则列表
	func (m *ruleManagerImpl) UpdateRules(rules []Rule) error {
		if rules == nil {
			return ErrInvalidRuleConfig
		}

		m.mux.Lock()
		defer m.mux.Unlock()

		m.rules = make([]Rule, len(rules))
		copy(m.rules, rules)
		logx.Infof("Updated %d rules", len(rules))
		return nil
	}

	// syncRules 定期同步规则
	func (m *ruleManagerImpl) syncRules(ctx context.Context) {
		ticker := time.NewTicker(time.Minute * 5)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-m.quit:
				return
			case <-ticker.C:
				if err := m.loadRulesFromStorage(); err != nil {
					logx.Errorf("Failed to sync rules: %v", err)
				}
			}
		}
	}

	// loadRulesFromStorage 从存储加载规则
	func (m *ruleManagerImpl) loadRulesFromStorage() error {
		// TODO: 实现从数据库或其他存储加载规则的逻辑
		// 这里需要添加实际的规则加载逻辑
		logx.Info("Loading rules from storage...")
		return nil
	}
