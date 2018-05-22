package inplace

import (
	"reflect"
	"testing"
)

func BenchmarkAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nonEmpty([]string{"", "a", "", "b", "", "c"})
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nonEmpty2([]string{"", "a", "", "b", "", "c"})
	}
}

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nonEmpty([]string{"", "a", "", "b", "", "c"})
	}
}

func TestNonEmpty(t *testing.T) {
	tests := []struct {
		in   []string
		want []string
	}{
		{
			[]string{""},
			[]string{},
		},
		{
			[]string{"a", "b", "c"},
			[]string{"a", "b", "c"},
		},
		{
			[]string{"1", "", "2", "3", ""},
			[]string{"1", "2", "3"},
		},
	}

	for _, test := range tests {
		got := nonEmpty3(test.in)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("nonEmpty3(%q)=%q, want:%q", test.in, got, test.want)
		}
	}

}
