package main

import (
	"fmt"
)

func main() {
	signTest(255, -127)
	signTest(255, -1)
	//	shiftInt(1)
	//	shiftInt(-96)
	//	shiftInt(-1)
}

func shiftInt(x int8) {
	fmt.Printf("Testing %d\n", x)

	for i := uint(0); i < 8; i++ {
		xs := x << i
		fmt.Printf("Shift: %d, signed: %d\n", i, xs)
		fmt.Printf("Signed: %08b\n\n", xs)
	}
}

// signTest tries to look at the bit representation signed integers
//
func signTest(u uint8, x int8) {

	fmt.Printf("Unsigned: %d, signed: %d\n", u, x)

	for i := uint(0); i < 8; i++ {
		us := u >> i
		xs := x >> i
		fmt.Printf("Shift: %d, Unsigned: %d, signed: %d\n", i, us, xs)
		fmt.Printf("Unsigned: %08b, signed: %08b\n\n", us, uint8(uint(xs)))
	}
}
