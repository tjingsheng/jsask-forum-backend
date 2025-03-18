package utils

func RemoveDuplicates[T string | int](strSlice []T) []T {
	allKeys := make(map[T]bool)
	result := []T{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			result = append(result, item)
		}
	}
	return result
}
