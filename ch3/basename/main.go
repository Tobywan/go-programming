package main

import (
	"strings"
)

func main() {

}
func basename(s string) string {
	s = s[strings.LastIndex(s, "/")+1:]

	if dotpos := strings.LastIndex(s, "."); dotpos >= 0 {
		s = s[:dotpos]
	}
	return s
}
