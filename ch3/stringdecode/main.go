package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	decode2("Toby")
	decode2("\U00004e16 \\ \u7123")
}

func decode(s string) {
	fmt.Println("--------------------")
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%d\t%c\n", i, size, r)
		i += size
	}

}

func decode2(s string) {
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)

	fmt.Println("--------------------")
	for i, r := range s {
		fmt.Printf("%d\t%c\n", i, r)

	}
}
