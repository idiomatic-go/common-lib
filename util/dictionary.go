package util

import "sync"

type DictionaryValue int32

const (
	DictionaryValueNotFound = DictionaryValue(-1)
)

type ReverseDictionary struct {
	threadSafe   bool
	m            map[string]DictionaryValue
	currentValue DictionaryValue
	mu           sync.RWMutex
}

func CreateReverseDictionary(threadSafe bool) *ReverseDictionary {
	return &ReverseDictionary{threadSafe: threadSafe, m: make(map[string]DictionaryValue, 1), currentValue: DictionaryValueNotFound}
}

func (d *ReverseDictionary) IsEmpty() bool {
	if d.threadSafe {
		d.mu.RLock()
		defer d.mu.RUnlock()
	}
	return d.currentValue == DictionaryValueNotFound
}

func (d *ReverseDictionary) Lookup(key string) DictionaryValue {
	if d.threadSafe {
		d.mu.RLock()
		defer d.mu.RUnlock()
	}
	if v, ok := d.m[key]; ok {
		return v
	}
	return DictionaryValueNotFound
}

func (d *ReverseDictionary) Add(key string) DictionaryValue {
	v := d.Lookup(key)
	if v != DictionaryValueNotFound {
		return v
	}
	if d.threadSafe {
		d.mu.Lock()
		defer d.mu.Unlock()
	}
	d.currentValue++
	d.m[key] = d.currentValue
	return d.currentValue
}
