package popcount

import ()

// pc is the population count of i
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountShift(x uint64) int {

	result := (x & 1)

	for shift := x / 2; shift > 0; shift /= 2 {
		result += (shift & 1)
	}
	return int(result)
}

func PopCountLoop(x uint64) int {
	// casting x to byte will take just the last 8 bits
	result := 0
	for i := uint(0); i < 8; i++ {
		result += int(pc[byte(x>>(i*8))])
	}
	return result
}

// PopCount returns the population count (number of bits in) x
func PopCountChunk(x uint64) int {
	// Divvy up the uint64 into 8 byte chunks
	// casting x to byte will take just the last 8 bits
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}
