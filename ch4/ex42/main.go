package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var alg string
var data string

func init() {
	flag.StringVar(&alg, "a", "256", "Which hashing algorithm? (256|384|512)")
	flag.StringVar(&data, "d", "", "Data to hash")
}

func sum256(in *[]byte) []byte {
	//result := make([]byte, sha256.Size)
	hash := sha256.Sum256(*in)
	return hash[:]
}

func sum384(in *[]byte) []byte {
	hash := sha512.Sum384(*in)
	return hash[:]
}

func sum512(in *[]byte) []byte {
	hash := sha512.Sum512(*in)
	return hash[:]
}

var hashfuncs = map[string]func(*[]byte) []byte{
	"256": sum256,
	"384": sum384,
	"512": sum512,
}

func main() {
	flag.Parse()
	f := hashfuncs[alg]
	if data == "" || f == nil {
		flag.Usage()
		os.Exit(1)
	}
	fmt.Printf("Hashing function: %q, data: %q\n", alg, data)
	bd := []byte(data)
	hash := f(&bd)
	fmt.Printf("%x\n", hash)
}
