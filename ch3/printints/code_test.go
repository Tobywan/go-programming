package printints

import "testing"

func TestIntsToString(t *testing.T) {
	tests := []struct {
		in   []int
		want string
	}{
		{[]int{}, "[]"},
		{[]int{0}, "[0]"},
		{[]int{1, 2}, "[1,2]"},
		{[]int{-9, 3, -27}, "[-9,3,-27]"},
	}
	for _, test := range tests {
		got := IntsToString(test.in)
		if got != test.want {
			t.Errorf("input=%v,got=%q, want=%q", test.in, got, test.want)
		}
	}
}
