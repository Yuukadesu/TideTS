package storageengine

import "time"

// Hooks 允许外层注入轻量观测回调，避免引擎直接依赖具体 metrics 框架。
type Hooks struct {
	OnWrite     func(op string, points int, duration time.Duration)
	OnRead      func(op string, points int, duration time.Duration)
	OnWAL       func(op string, records int)
	OnTombstone func(op string, ranges int)
	OnFlush     func(points int, duration time.Duration)
	OnCompact   func(duration time.Duration, inputFiles, outputFiles int)
}

// SetHooks 设置引擎观测 hooks。
func (e *Engine) SetHooks(h Hooks) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.hooks = h
}
