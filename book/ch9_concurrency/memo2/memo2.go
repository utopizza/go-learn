package main

import "sync"

// Memo 缓存
type Memo struct {
	f     Func
	mux   sync.Mutex
	cache map[string]Result
}

// Func 需要被缓存的函数
type Func func(key string) (interface{}, error)

// Result 函数的执行结果
type Result struct {
	value interface{}
	err   error
}

// NewMemo 创建缓存
func NewMemo(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]Result),
	}
}

// Get 查询缓存
func (m *Memo) Get(key string) (interface{}, error) {
	m.mux.Lock() // 解决 data race 问题
	res, ok := m.cache[key]
	if !ok { // 没有命中缓存，则执行f，并缓存
		res.value, res.err = m.f(key)
		m.cache[key] = res
	}
	m.mux.Unlock()
	return res.value, res.err
}
