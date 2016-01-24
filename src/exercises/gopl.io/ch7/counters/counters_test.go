package counters

import (
	"bytes"
	"fmt"
	"io"
)

func Example1() {
	var c WordCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c.Write([]byte("goodbye cruel world"))
	fmt.Println(c)

	c.Write(nil)
	fmt.Println(c)
	// Output:
	// 1
	// 4
	// 4
}

func Example2() {
	var c LineCounter
	c.Write([]byte("hello\n"))
	fmt.Println(c)

	c.Write([]byte("goodbye\ncruel\nworld"))
	fmt.Println(c)

	c.Write(nil)
	fmt.Println(c)
	// Output:
	// 1
	// 4
	// 4
}

func Example3() {
	var b bytes.Buffer
	var w io.Writer
	var cntr *int64
	w, cntr = CountingWriter(&b)
	n, err := w.Write([]byte("Hello"))
	if err != nil {
		panic(fmt.Sprintf("Failed to write: %v", err))
	}

	fmt.Println(b.String())
	fmt.Println(n)
	fmt.Println(*cntr)

	n, err = w.Write([]byte(" world"))
	if err != nil {
		panic(fmt.Sprintf("Failed to write: %v", err))
	}
	fmt.Println(b.String())
	fmt.Println(n)
	fmt.Println(*cntr)

	// Output:
	// Hello
	// 5
	// 5
	// Hello world
	// 6
	// 11
}
