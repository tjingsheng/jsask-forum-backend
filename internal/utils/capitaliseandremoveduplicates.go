package utils

func CapitaliseAndRemoveDuplicates(str []string) []string {
	var result []string
	for _, item := range str {
		result = append(result, CapitaliseWords(item))
	}
	result = RemoveDuplicates(result)
	return result
}
