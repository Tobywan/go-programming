package main

import (
	"fmt"
	"github.com/Tobywan/go-programming/ch2/popcount"
)

func main() {

	test(uint64(1.75643326e18 + 7))
	test(255)
	test(0)
	test(65536)
}

func test(x uint64) {
	fmt.Printf("Testing: %d\n", x)
	fmt.Printf("PopCountLoop:%d\n", popcount.PopCountLoop(x))
	fmt.Printf("PopCountChunk:%d\n", popcount.PopCountChunk(x))
	fmt.Printf("PopCountShift:%d\n", popcount.PopCountShift(x))
}
