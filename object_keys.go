package utils

func ObjectKeys(data map[string]interface{}) []string {
	var keys []string
	for v := range data {
		keys = append(keys, v)
	}
	return keys
}
