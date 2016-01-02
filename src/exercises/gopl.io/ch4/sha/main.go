package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"strings"
)

func main() {
	var shaVer = flag.String("s", "sha256", "Type of sha to use")
	flag.Parse()

	var f func(data []byte) string
	switch strings.ToLower(*shaVer) {
	case "sha256":
		f = func(data []byte) string {
			return fmt.Sprintf("%x", sha256.Sum256(data))
		}
	case "sha384":
		f = func(data []byte) string {
			return fmt.Sprintf("%x", sha512.Sum384(data))
		}
	case "sha512":
		f = func(data []byte) string {
			return fmt.Sprintf("%x", sha512.Sum512(data))
		}
	default:
		fmt.Printf("Unknown sha type %s\n", *shaVer)
		return
	}

	var args = flag.Args()
	for i := 0; i < len(args); i++ {
		arg := []byte(args[i])
		fmt.Printf("%x\n", f(arg))
	}
}
