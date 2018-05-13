package comma

import (
	"fmt"
)

// Insert places a comma every 3 chars in an integer for both positive and
// negative
func Insert(n int) string {
	s := fmt.Sprintf("%d", n)
	return insertComma(s)
}

func insertComma(s string) string {
	l := len(s)
	if (l < 4) || (l == 4 && s[0] == byte('-')) {
		// too short, or negative prefix, don't process
		return s
	}

	return insertComma(s[:l-3]) + "," + s[l-3:]

}
