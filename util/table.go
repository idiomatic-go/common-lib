package util

import "sync"

type IsEqual func(key, val any) bool

type AnyTable struct {
	data  []any
	equal IsEqual
	mu    sync.RWMutex
}

func CreateAnyTable(fn IsEqual) *AnyTable {
	if fn == nil {
		fn = func(key any, val any) bool { return false }
	}
	return &AnyTable{data: nil, equal: fn}
}

func (t *AnyTable) Empty() {
	if t == nil {
		return
	}
	t.mu.Lock()
	t.data = nil
	t.mu.Unlock()
}

func (t *AnyTable) Count() int {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return len(t.data)
}

func (t *AnyTable) Delete(key any) bool {
	if t == nil || key == nil {
		return false
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	if i, ok := t.find(key); ok {
		t.data[i] = t.data[len(t.data)-1] // Copy last element to index i.
		t.data[len(t.data)-1] = nil       // Erase last element (write zero value).
		t.data = t.data[:len(t.data)-1]   // Truncate slice.
		return true
	}
	return false
}

func (t *AnyTable) Put(val any) bool {
	if t == nil || val == nil {
		return false
	}
	t.mu.Lock()
	defer t.mu.Unlock()
	t.data = append(t.data, val)
	return true
}

func (t *AnyTable) Get(key any) (any, bool) {
	if t == nil || key == nil {
		return nil, false
	}
	t.mu.RLock()
	defer t.mu.RUnlock()
	if i, ok := t.find(key); ok {
		return t.data[i], ok
	}
	return nil, false
}

func (t *AnyTable) find(key any) (int, bool) {
	if t == nil || key == nil {
		return -1, false
	}
	for i, _ := range t.data {
		if t.equal(key, t.data[i]) {
			return i, true
		}
	}
	return -1, false
}
