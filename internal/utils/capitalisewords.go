package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func CapitaliseWords(str string) string {
	return cases.Title(language.Und).String(str)
}
