package inplace

import (
	"reflect"
	"testing"
)

func BenchmarkAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nonEmpty1([]string{"", "a", "", "b", "", "c"})
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nonEmpty2([]string{"", "a", "", "b", "", "c"})
	}
}

func BenchmarkCopy(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nonEmpty3([]string{"", "a", "", "b", "", "c"})
	}
}
func TestNonEmpty1(t *testing.T) {
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
		got := nonEmpty1(test.in)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("nonEmpty3(%q)=%q, want:%q", test.in, got, test.want)
		}
	}

}
func TestNonEmpty2(t *testing.T) {
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
		got := nonEmpty2(test.in)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("nonEmpty3(%q)=%q, want:%q", test.in, got, test.want)
		}
	}

}

func TestNonEmpty3(t *testing.T) {
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

func TestReverse(t *testing.T) {
	tests := []struct {
		in   []string
		want []string
	}{
		{
			[]string{""},
			[]string{""},
		},
		{
			[]string{"a"},
			[]string{"a"},
		},
		{
			[]string{"5", "4", "3", "2", "1"},
			[]string{"1", "2", "3", "4", "5"},
		},
		{
			[]string{"6", "5", "4", "3", "2", "1"},
			[]string{"1", "2", "3", "4", "5", "6"},
		},
	}

	for _, test := range tests {

		got := reverse(test.in)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("reverse(%q)=%q, want:%q", test.in, got, test.want)
		}
	}

}
func TestRotate(t *testing.T) {
	tests := []struct {
		in   []string
		n    int
		want []string
	}{
		{
			[]string{""},
			8,
			[]string{""},
		},
		{
			[]string{"a"},
			0,
			[]string{"a"},
		},
		{
			[]string{"1", "2", "3", "4", "5"},
			1,
			[]string{"2", "3", "4", "5", "1"},
		},
		{
			[]string{"1", "2", "3", "4", "5"},
			2,
			[]string{"3", "4", "5", "1", "2"},
		},
		{
			[]string{"1", "2", "3", "4", "5"},
			3,
			[]string{"4", "5", "1", "2", "3"},
		},
		{
			[]string{"1", "2", "3", "4", "5"},
			4,
			[]string{"5", "1", "2", "3", "4"},
		},
		{
			[]string{"1", "2", "3", "4", "5"},
			5,
			[]string{"1", "2", "3", "4", "5"},
		},
		{
			[]string{"1", "2", "3", "4", "5"},
			6,
			[]string{"2", "3", "4", "5", "1"},
		},
		{
			[]string{"1", "2", "3", "4", "5"},
			7,
			[]string{"3", "4", "5", "1", "2"},
		},
		{
			[]string{"1", "2", "3", "4", "5"},
			8,
			[]string{"4", "5", "1", "2", "3"},
		},
		{
			[]string{"1", "2", "3", "4", "5"},
			9,
			[]string{"5", "1", "2", "3", "4"},
		},
		{
			[]string{"1", "2", "3", "4", "5"},
			10,
			[]string{"1", "2", "3", "4", "5"},
		},
	}

	for _, test := range tests {

		got := rotate(test.in, test.n)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("rotate(%q,%d)=%q, want:%q", test.in, test.n, got, test.want)
		}
	}

}
