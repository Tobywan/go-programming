package inplace

import (
	"unicode"
	"unicode/utf8"
)

// nonEmpty1 modifies the slice inplace to squash empty strings
func nonEmpty1(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// nonEmpty2 uses append to a second slice pointing at the same underlying array
func nonEmpty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// nonEmpty3 uses copy to close up on the empty strings
// but range iterator is invalid
func nonEmpty3(strings []string) []string {
	n := len(strings)

	for i := 0; i < n; i++ {
		s := strings[i]
		if s == "" {
			copy(strings[i:], strings[i+1:])
			n--
		}
	}
	return strings[:n]
}

// reverse does an in place reversal of a slice
func reverse(s []string) []string {
	l := len(s)
	if l <= 1 {
		return s
	}
	for i, j := 0, l-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// rotate shift the elements to the left by n and puts the over flowed ones back on the right
func rotate(s []string, n int) []string {
	l := len(s)
	if l <= 1 {
		return s
	}

	m := n % l // Deal with n  > l
	if m == 0 {
		return s
	}

	r := s[:0]
	r = append(r, reverse(s[0:m])...)
	r = append(r, reverse(s[m:])...)
	r = reverse(r)
	return r
}

// rotateOnce shift the elements to the left by n and puts the over flowed ones back on the right
// exercise 4.4
func rotateOnce(s []string, n int) []string {
	l := len(s)
	if l <= 1 {
		return s
	}

	m := n % l // Deal with n  > l
	if m == 0 {
		return s
	}
	// Append the beginning to the end
	s = append(s, s[:m]...)
	return s[m:]
}

// dedupe removes adjacent dupicates in a string slice
// EX 4.5
func dedupe(strings []string) []string {

	if len(strings) <= 1 {
		// no duplicates
		return strings
	}
	last := strings[0]
	i := 1
	for _, s := range strings[1:] {
		if s != last {
			strings[i] = s
			last = s
			i++
		}
	}
	return strings[:i]
}

// asciispaceRange converts runs of unicode spaces in a byte slice into a single
// ascii space
// EX 4.6
func asciispaceRange(b []byte) []byte {
	// consider the bytes as a range of runes
	ret := make([]byte, len(b))
	pos := 0
	spaces := false // true for consecutive spaces
	for _, r := range string(b) {
		if unicode.IsSpace(r) {
			if spaces {
				continue
			}
			spaces = true
			r = ' '
		} else {
			spaces = false
		}
		n := utf8.EncodeRune(ret[pos:], r)
		pos += n
	}

	return ret[:pos]

}

// asciispaceInplace converts runs of unicode spaces in a byte slice into a single
// ascii space
// EX 4.6
func asciispaceInplace(b []byte) []byte {
	// consider the bytes as a range of runes
	l := len(b)

	spaces := false // true for consecutive spaces
	posOut := 0
	for posIn := 0; posIn < l; {
		r, n := utf8.DecodeRune(b[posIn:])
		posIn += n
		if unicode.IsSpace(r) {
			if spaces {
				continue
			}
			spaces = true
			r = ' '
		} else {
			spaces = false
		}
		m := utf8.EncodeRune(b[posOut:], r)
		posOut += m
	}

	return b[:posOut]

}
