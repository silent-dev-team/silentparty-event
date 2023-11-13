package utils

import "strings"

func ParseArrayString(s string) []string {
	var result []string
	s = strings.Trim(s, "[]") // trim brackets
	// split by comma
	parts := strings.Split(s, ",")

	for _, part := range parts {
		trim := part
		trim = strings.Trim(trim, " ")
		trim = strings.Trim(trim, "\"")
		trim = strings.Trim(trim, " ")
		result = append(result, trim)
	}
	return result
}
