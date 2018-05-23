package inplace

import (
	"reflect"
	"testing"
	"unicode"
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

func BenchmarkRotate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotate([]string{"1", "2", "3", "4", "5"}, 4)
	}
}

func BenchmarkRotateOnce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		rotateOnce([]string{"1", "2", "3", "4", "5"}, 4)
	}
}

func TestRotateOnce(t *testing.T) {
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

		got := rotateOnce(test.in, test.n)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("rotate(%q,%d)=%q, want:%q", test.in, test.n, got, test.want)
		}
	}

}

func TestDeDupe1(t *testing.T) {
	tests := []struct {
		in   []string
		want []string
	}{
		{
			[]string{"a"},
			[]string{"a"},
		},
		{
			[]string{},
			[]string{},
		},
		{
			[]string{"a", "b"},
			[]string{"a", "b"},
		},
		{
			[]string{"a", "b", "b"},
			[]string{"a", "b"},
		},
		{
			[]string{"1", "2", "2", "3", "3"},
			[]string{"1", "2", "3"},
		},
		{
			[]string{"1", "1", "1", "1", "1"},
			[]string{"1"},
		},
		{
			[]string{"1", "1", "1", "1", "3"},
			[]string{"1", "3"},
		},
	}

	for _, test := range tests {
		got := dedupe(test.in)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("dedupe(%q)=%q, want:%q", test.in, got, test.want)
		}
	}

}

func TestIsSpace(t *testing.T) {
	s := "\u0020\u00A0\u1680\u2000\u2001\u2002\u2003\u2004\u2005\u2006\u2007\u2008\u2009\u200A\u202F\u205F\u3000"
	for _, r := range s {
		t.Logf("%U - %t", r, unicode.IsSpace(r))
	}
}

func BenchmarkAsciispaceRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := []byte("X\u0020\u00A0\u1680\u2000\u2001\u2002X\u2003X\u2004\u2005\u2006\u2007\u2008\u2009\u200A\u202F\u205FX\u3000X")
		asciispaceRange(in)
	}
}
func BenchmarkAsciispaceInplace(b *testing.B) {
	for i := 0; i < b.N; i++ {
		in := []byte("X\u0020\u00A0\u1680\u2000\u2001\u2002X\u2003X\u2004\u2005\u2006\u2007\u2008\u2009\u200A\u202F\u205FX\u3000X")
		asciispaceInplace(in)
	}
}
func TestAsciispaceInplace(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{
			in:   "",
			want: "",
		},
		{
			in:   "         ",
			want: " ",
		},
		{
			in:   "a",
			want: "a",
		},
		{
			in:   " a ",
			want: " a ",
		},
		{
			in:   "  a  b  c  d  ",
			want: " a b c d ",
		},
		{
			in:   " a ",
			want: " a ",
		},
		{
			in:   "X\u0020\u00A0\u1680\u2000\u2001\u2002X\u2003X\u2004\u2005\u2006\u2007\u2008\u2009\u200A\u202F\u205FX\u3000X",
			want: "X X X X X",
		},
	}

	for _, test := range tests {
		bIn := []byte(test.in)
		bWant := []byte(test.want)

		got := asciispaceInplace(bIn)

		if !reflect.DeepEqual(bWant, got) {
			t.Errorf("asciispaceInplace(%v)=%v, want %v", bIn, got, bWant)
		}
	}

}
