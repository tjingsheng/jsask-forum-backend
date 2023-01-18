package utils

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	result := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			result = append(result, item)
		}
	}
	return result
}
