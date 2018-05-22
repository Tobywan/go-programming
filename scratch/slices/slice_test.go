package slices

import (
	"fmt"
	"testing"
)

func TestCreate(t *testing.T) {
	// Our conventional way to create slice
	a := make([]int, 4, 5)
	// Or create the array first
	b := [4]int{}
	// And wrap it with a zero length slice
	c := b[:0]
	// or a slice with a longer length
	d := c[:4]
	// Would panic as cap is exceeded:
	// e := a[:6]
	// t.Logf("%v", e)

	t.Logf("a: %v, len: %d cap: %d", a, len(a), cap(a))
	t.Logf("b: %v, len: %d", b, len(b))
	t.Logf("c: %v, len: %d cap: %d", c, len(c), cap(c))
	t.Logf("d: %v, len: %d cap: %d", d, len(d), cap(d))

	c = c[:3]
	a[0] = 1
	// b, c, d all refer to the same unrelying array
	b[1] = 2
	c[2] = 3
	d[3] = 4
	t.Logf("a: %v, len: %d cap: %d", a, len(a), cap(a))
	t.Logf("b: %v, len: %d", b, len(b))
	t.Logf("c: %v, len: %d cap: %d", c, len(c), cap(c))
	t.Logf("d: %v, len: %d cap: %d", d, len(d), cap(d))
}

func TestAppend(t *testing.T) {
	dst := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	src := []int{10, 11, 12}
	grow := make([]int, len(dst)+len(src))

	t.Logf(sniff(dst, "dst"))
	t.Logf(sniff(src, "src"))
	t.Logf(sniff(grow, "grow"))

	n := copy(grow[:len(dst)], dst)
	t.Logf("copied: %d", n)
	t.Logf(sniff(grow, "grow"))

	n = copy(grow[len(dst):], src)
	t.Logf("copied: %d", n)
	t.Logf(sniff(grow, "grow"))
}

func sniff(sl []int, name string) string {
	return fmt.Sprintf("%s:%v len:%d cap:%d", name, sl, len(sl), cap(sl))
}
