package printints

import (
	"bytes"
	"fmt"
)

// IntsToString formats a slice of ints and puts , between them
func IntsToString(values []int) string {

	var buf bytes.Buffer

	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteByte(',')
		}
		if _, err := fmt.Fprintf(&buf, "%d", v); err != nil {
			continue
		}
	}
	buf.WriteByte(']')
	return buf.String()
}
