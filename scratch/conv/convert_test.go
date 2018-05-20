package conv

import (
	"testing"
)

// Does not compile
// func TestBigINtToBytes(t *testing.T) {
// 	n := 12345678901234567890
// 	b := []byte(n)
// 	t.Logf("b: %v", b)

// }
func TestCompile(t *testing.T) {
	var i = 1
	var i32 int32 = 2
	var i64 int64 = 3

	res := i + int(i32) + int(i64)

	t.Logf("%d + %d + %d = %d", i, i32, i64, res)

	x := 2
	x32 := int32(4)
	x64 := int64(8)
	res = x + int(x32) + int(x64)

	t.Logf("%d + %d + %d = %d", x, x32, x64, res)

}

func TestCast(t *testing.T) {
	var f float64

	f = 5 / 9
	t.Logf("f=%f", f)

	f = 5.0 / 9
	t.Logf("f=%f", f)

	f = 5 / 9.0
	t.Logf("f=%f", f)

	f = (5 + 9) / 3
	t.Logf("f=%f", f)

	f = 2 + 0i
	t.Logf("f=%f", f)

}
