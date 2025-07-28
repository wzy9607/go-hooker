package hooker

import (
	"slices"
	"sync"
)

type Hook[Func any] func(next Func) Func

type Hooker[Func any] struct {
	mu sync.RWMutex

	origin  Func         // 原始函数
	wrapped Func         // 包装后的函数
	hooks   []Hook[Func] // 钩子函数列表
}

func NewHooker[Func any](origin Func) *Hooker[Func] {
	h := &Hooker[Func]{
		mu:      sync.RWMutex{},
		origin:  origin,
		wrapped: origin,
		hooks:   make([]Hook[Func], 0),
	}
	return h
}

func (h *Hooker[Func]) chain() {
	h.wrapped = h.origin
	for i := len(h.hooks) - 1; i >= 0; i-- {
		h.wrapped = h.hooks[i](h.wrapped)
	}
}

func (h *Hooker[Func]) AddHook(hook ...Hook[Func]) *Hooker[Func] {
	h.mu.Lock()
	defer h.mu.Unlock()

	h.hooks = append(h.hooks, hook...)
	h.chain()
	return h
}

func (h *Hooker[Func]) GetOrigin() Func {
	h.mu.RLock()
	defer h.mu.RUnlock()

	return h.origin
}

func (h *Hooker[Func]) GetWrapped() Func {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return h.wrapped
}

func (h *Hooker[Func]) GetHooks() []Hook[Func] {
	h.mu.RLock()
	defer h.mu.RUnlock()
	return slices.Clone(h.hooks)
}
