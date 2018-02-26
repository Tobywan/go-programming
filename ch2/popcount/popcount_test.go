package popcount

import (
	"testing"
)

var data = []uint64{uint64(1.75643326e18 + 7),
	255,
	0,
	65536}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(data[0])
	}
}

func BenchmarkPopCountChunk(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountChunk(data[0])
	}
}

func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(data[0])
	}
}

/*
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
*/
