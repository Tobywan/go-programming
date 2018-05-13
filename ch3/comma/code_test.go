package comma

import "testing"

func TestInsert(t *testing.T) {
	tests := []struct {
		input int
		want  string
	}{
		{0, "0"},
		{100, "100"},
		{1234, "1,234"},
		{-10000, "-10,000"},
		{-210000, "-210,000"},
		{999999999999, "999,999,999,999"},
	}

	for _, test := range tests {
		got := Insert(test.input)
		if got != test.want {
			t.Errorf("Insert(%d)=%q, want %q", test.input, got, test.want)
		}
	}
}
