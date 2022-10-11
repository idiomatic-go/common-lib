package util

type DictionaryValue int32

const (
	DictionaryValueNotFound = DictionaryValue(-1)
)

// InvertedDictionary - used for inverted content, where key = string, value = id.
////                    The im is used for normal lookups key = id, value = string
type InvertedDictionary struct {
	m            map[string]DictionaryValue
	im           map[DictionaryValue]string
	currentValue DictionaryValue
}

func CreateInvertedDictionary() *InvertedDictionary {
	return &InvertedDictionary{m: make(map[string]DictionaryValue, 1), im: nil, currentValue: DictionaryValueNotFound}
}

func (d *InvertedDictionary) IsEmpty() bool {
	return d.currentValue == DictionaryValueNotFound
}

func (d *InvertedDictionary) Lookup(key string) DictionaryValue {
	if v, ok := d.m[key]; ok {
		return v
	}
	return DictionaryValueNotFound
}

func (d *InvertedDictionary) InverseLookup(value DictionaryValue) (string, bool) {
	if d.im == nil {
		d.im = make(map[DictionaryValue]string, 1)
		for k, v := range d.m {
			d.im[v] = k
		}
	}
	if key, ok := d.im[value]; ok {
		return key, true
	}
	return "", false
}

func (d *InvertedDictionary) Add(key string) DictionaryValue {
	v := d.Lookup(key)
	if v != DictionaryValueNotFound {
		return v
	}
	d.currentValue++
	d.m[key] = d.currentValue
	return d.currentValue
}
