package utils

import "fmt"

func CapitaliseAndRemoveDuplicates(str []string) []string {
	var result []string
	for _, item := range str {
		result = append(result, CapitaliseWords(item))
	}
	result = RemoveDuplicateStr(result)
	fmt.Print(result)
	return result
}
