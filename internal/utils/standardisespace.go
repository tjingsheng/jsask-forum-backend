package utils

import (
	"strings"
)

func StandardiseSpace(str string) string {
	return strings.Join(strings.Fields(str), " ")
}
