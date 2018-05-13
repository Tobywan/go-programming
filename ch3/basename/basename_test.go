package ch3

import (
	"testing"
)

func TestBasename(t *testing.T) {
	tests := []struct {
		path string
		want string
	}{
		{"a/b/c.go", "c"},
		{"c.d.go", "c.d"},
		{"abc", "abc"},
		{".test", ""},
		{"", ""},
	}
	for _, test := range tests {
		got := Basename(test.path)
		if got != test.want {
			t.Errorf("basename(%s)=%s, want %s", test.path, got, test.want)
		}
	}
}
