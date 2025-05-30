package site

import "strings"

func mapQuotes(str string) string {
	return "\"" + str + "\""
}

func unmapQuotes(str string) string {
	return strings.TrimSuffix(strings.TrimPrefix(str, "\""), "\"")
}
