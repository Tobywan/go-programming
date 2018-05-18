package iota

import (
	"math/big"
	"testing"
)

func TestKiBi(t *testing.T) {
	y := big.NewInt(0)
	y.SetString("1208925819614629174706176", 10)

	t.Logf("%s: %s\n", "yiB", y)
}
