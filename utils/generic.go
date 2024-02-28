package utils

func ArrayToBoolMap[TKey comparable](keys []TKey) map[TKey]bool {
	result := make(map[TKey]bool, len(keys))
	for _, key := range keys {
		result[key] = true
	}
	return result
}
