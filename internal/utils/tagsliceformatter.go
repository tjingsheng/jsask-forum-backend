package utils

// StandardiseSpace, CapitaliseWords,, and RemoveDuplicatester
func TagSliceFormatter(str []string) []string {
	var result []string
	for _, item := range str {
		result = append(result, CapitaliseWords(StandardiseSpace(item)))
	}
	result = RemoveDuplicates(result)
	return result
}
