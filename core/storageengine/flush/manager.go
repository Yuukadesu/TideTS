package flush

import (
	"sync"

	"github.com/hanami/tidets/commons/errors"
)

// Manager 异步 flush 工作池（单 worker，保证落盘顺序）。
type Manager struct {
	ch     chan func()
	wg     sync.WaitGroup
	closed bool
	mu     sync.Mutex
}

func NewManager(queue int) *Manager {
	if queue <= 0 {
		queue = 64
	}
	m := &Manager{ch: make(chan func(), queue)}
	m.wg.Add(1)
	go func() {
		defer m.wg.Done()
		for fn := range m.ch {
			fn()
		}
	}()
	return m
}

func (m *Manager) Submit(fn func()) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.closed {
		return commons.ErrFlushManagerClosed
	}
	select {
	case m.ch <- fn:
		return nil
	default:
		return commons.ErrFlushQueueFull
	}
}

func (m *Manager) Close() {
	m.mu.Lock()
	if m.closed {
		m.mu.Unlock()
		return
	}
	m.closed = true
	m.mu.Unlock()
	close(m.ch)
	m.wg.Wait()
}
