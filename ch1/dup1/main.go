// Counts duplicates on the stdin
// and writes them to stdout

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create a map to store the results
	counts := make(map[string]int)

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
