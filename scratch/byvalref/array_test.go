package byvalref

import (
	"testing"
)

func TestPassVal(t *testing.T) {
	a := [...]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	modifyVal(a)
	t.Logf("%v", a)
	modifyRef(&a)
	t.Logf("%v", a)

}

func modifyVal(a [26]rune) {
	a[0] = '@'
}

func modifyRef(a *[26]rune) {
	a[0] = '@'
}
