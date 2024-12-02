package utils

import (
	"strings"
)

func GenerateSlug(input string) string {
	return strings.ToLower(strings.ReplaceAll(input, " ", "-"))
}
