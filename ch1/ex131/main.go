// Counts duplicates on the stdin
// and writes them to stdout
// or takes a list of files as arguments

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// create a map to store the results
	counts := make(map[string]map[string]bool)

	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2 : %v\n", err)
			continue
		}
		countLines(f, counts, arg)
		f.Close()
	}
	for line, occ := range counts {
		if len(occ) > 1 {
			fmt.Printf("%d\t%s\t%v\n", len(occ), line, occ)
		}
	}
}

// Count the lines of the file/stream passed in
func countLines(f *os.File, counts map[string]map[string]bool, fname string) {
	input := bufio.NewScanner(f)

	for input.Scan() {
		s := input.Text()
		occ := counts[s]
		if len(occ) == 0 {
			occ = make(map[string]bool)
			counts[s] = occ
		}
		occ[fname] = true

	}

}
