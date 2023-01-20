package utils

// StandardiseSpace, CapitaliseWords
func UsernameFormatter(str string) string {
	return CapitaliseWords(StandardiseSpace(str))
}
