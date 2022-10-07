package util

type IsEqual func(key, val any) bool

type AnyTable struct {
	data  []any
	equal IsEqual
}

func CreateAnyTable(fn IsEqual) *AnyTable {
	if fn == nil {
		fn = func(key any, val any) bool { return false }
	}
	return &AnyTable{nil, fn}
}

func (t *AnyTable) Empty() {
	t.data = nil
}

func (t *AnyTable) Count() int {
	return len(t.data)
}

func (t *AnyTable) Delete(key any) bool {
	if t == nil || key == nil {
		return false
	}
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
	t.data = append(t.data, val)
	return true
}

func (t *AnyTable) Get(key any) (any, bool) {
	if t == nil || key == nil {
		return nil, false
	}
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
