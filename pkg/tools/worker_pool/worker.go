package worker_pool

import (
	"sync"

	"github.com/ape902/corex/logx"
)

// WorkerPool 工作池结构体
type WorkerPool struct {
	size    int         // 工作者数量
	tasks   chan func() // 任务队列
	workers []*worker   // 工作者列表
	wg      sync.WaitGroup
	quit    chan struct{} // 关闭信号
}

// worker 工作者结构体
type worker struct {
	pool *WorkerPool
}

// NewWorkerPool 创建新的工作池
func NewWorkerPool(size int) *WorkerPool {
	if size <= 0 {
		size = 10 // 默认工作者数量
	}

	return &WorkerPool{
		size:  size,
		tasks: make(chan func(), 1000), // 任务队列容量
		quit:  make(chan struct{}),
	}
}

// Start 启动工作池
func (p *WorkerPool) Start() {
	logx.Info("Starting worker pool...")

	// 创建工作者
	p.workers = make([]*worker, p.size)
	for i := 0; i < p.size; i++ {
		w := &worker{pool: p}
		p.workers[i] = w
		p.wg.Add(1)
		go w.run()
	}
}

// Stop 停止工作池
func (p *WorkerPool) Stop() {
	logx.Info("Stopping worker pool...")
	close(p.quit)
	p.wg.Wait()
	close(p.tasks)
}

// Submit 提交任务到工作池
func (p *WorkerPool) Submit(task func()) {
	select {
	case p.tasks <- task:
		// 任务提交成功
	case <-p.quit:
		// 工作池已关闭
		logx.Warn("Worker pool is shutting down, task rejected")
	}
}

// run 工作者运行循环
func (w *worker) run() {
	defer w.pool.wg.Done()

	for {
		select {
		case task := <-w.pool.tasks:
			// 执行任务
			func() {
				defer func() {
					if r := recover(); r != nil {
						logx.Errorf("Panic in worker task: %v", r)
					}
				}()
				task()
			}()
		case <-w.pool.quit:
			// 收到退出信号
			return
		}
	}
}
