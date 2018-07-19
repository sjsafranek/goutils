package utils

func ObjectKeys(data map[string]interface) []string {
	var keys []string
	for _, v := range params {
		keys = append(keys, v)
	}
	return keys
}
