package counters

import "fmt"

func Example1() {
	var c WordCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c.Write([]byte("goodbye cruel world"))
	fmt.Println(c)

	c.Write(nil)
	fmt.Println(c)
	// Output
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
	// Output
	// 1
	// 4
	// 4
}
