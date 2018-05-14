package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

var hashfuncs = map[string]func(){
	"256": sha256.Sum256,
	"384": sha384.Sum384,
	"512": shah512.Sum512,
}

func main() {

}
