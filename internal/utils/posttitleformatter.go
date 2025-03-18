package utils

// StandardiseSpace, CapitaliseWords
func PostTitleFormatter(str string) string {
	return CapitaliseWords(StandardiseSpace(str))
}
