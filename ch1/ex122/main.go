// Uses range to iterate over arguments

package main

import (
	"fmt"
	"os"
)

func main() {
	for n, arg := range os.Args[1:] {
		fmt.Printf("Index:%d, Argument:%s\n", n, arg)
	}
}
