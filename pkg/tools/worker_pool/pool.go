package worker_pool

// Pool 是WorkerPool的别名，用于向后兼容
type Pool = WorkerPool

// NewPool 创建新的工作池，是NewWorkerPool的别名
func NewPool(size int) *Pool {
	pool := NewWorkerPool(size)
	pool.Start() // 自动启动工作池
	return pool
}