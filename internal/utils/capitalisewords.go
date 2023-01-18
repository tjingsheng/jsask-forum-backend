package utils

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"strings"
)

func CapitaliseWords(str string) string {
	return strings.TrimSpace(cases.Title(language.Und).String(str))
}
