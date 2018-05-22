package inplace

// nonEmpty modifies the slice inplace to squash empty strings
func nonEmpty(strings []string) []string {
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
