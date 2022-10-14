package util

import "sync"

type List map[any]struct{}

func (l List) IsEmpty() bool {
	if l == nil {
		return true
	}
	return len(l) == 0
}

func (l List) Contains(item any) bool {
	if l == nil {
		return false
	}
	_, ok := l[item]
	return ok
}

func (l List) Add(item any) {
	if l == nil {
		return
	}
	l[item] = struct{}{}
}

func (l List) Remove(item any) {
	if l == nil {
		return
	}
	delete(l, item)
}

type SyncList sync.Map

func (l *SyncList) Contains(item any) bool {
	if l == nil {
		return false
	}
	m := sync.Map(*l)
	_, ok := m.Load(item)
	return ok
}

func (l *SyncList) Add(item any) {
	if l == nil {
		return
	}
	m := sync.Map(*l)
	m.Store(item, struct{}{})
	_, ok := m.Load(item)
	if ok {
	}
}

func (l *SyncList) Remove(item any) {
	if l == nil {
		return
	}
	m := sync.Map(*l)
	m.Delete(item)
}
