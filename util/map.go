package util

func MapSubset(keys []string, m map[string]string) map[string]string {
	var result map[string]string

	if m == nil || keys == nil {
		return nil
	}
	for _, key := range keys {
		if value, ok := m[key]; ok {
			if result == nil {
				result = make(map[string]string)
			}
			result[key] = value
		}
	}
	return result
}
