package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Replace each substring $<value> in s with
//  the results of f("value")
func expand(s string, f func(string) string) string {
	var accumulatingVar bool
	var buffer bytes.Buffer
	var currVar bytes.Buffer
	accumulatingVar = false
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		if accumulatingVar {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				currVar.WriteRune(r)
			} else {
				// Flush variable, then process new one
				accumulatingVar = false
				buffer.WriteString(f(currVar.String()))
				currVar.Reset()

				if r == '$' {
					// Started a new one -- e.g. 'this is $a$string$to interpolate'
					accumulatingVar = true
				} else {
					buffer.WriteRune(r)
				}
			}
		} else {
			if r == '$' {
				accumulatingVar = true
			} else {
				buffer.WriteRune(r)
			}
		}
		i += size
	}
	// If we ended still in accumulating mode, we need to flush
	if accumulatingVar {
		if currVar.Len() > 0 {
			buffer.WriteString(f(currVar.String()))
		} else {
			// Ended with a $
			buffer.WriteByte('$')
		}
	}
	return buffer.String()
}

func main() {
	testString := "This is a test $foo string ending in a $"
	testFunc := func(s string) string { return strings.ToUpper(s) }
	retVal := expand(testString, testFunc)
	fmt.Printf("Expanded string %s\n", retVal)
}
