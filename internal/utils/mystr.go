package mystr

import "strings"

// ToLowerKebab converts strings to kebab case with all lowercase characcters
func ToLowerKebab(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "-")
	return s
}