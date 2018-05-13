package ch3

import (
	"strings"
)

// Basename finds the last part of a path, minus any extension
func Basename(s string) string {
	s = s[strings.LastIndex(s, "/")+1:]

	if dotpos := strings.LastIndex(s, "."); dotpos >= 0 {
		s = s[:dotpos]
	}
	return s
}
