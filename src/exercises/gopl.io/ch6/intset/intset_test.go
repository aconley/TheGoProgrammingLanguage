package intset

import "fmt"

func Example1() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(x.Len())    // "3"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"
	fmt.Println(y.Len())    // "2"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x.Len())    // "4"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	// Output:
	// {1 9 144}
	// 3
	// {9 42}
	// 2
	// {1 9 42 144}
	// 4
	// true false
}

func Example2() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x.Len())    // "4"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"

	x.Remove(9)
	fmt.Println(x.String()) // "{1 42 144}"
	fmt.Println(x.Len())    // 3

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// 4
	// {[4398046511618 0 65536]}
	// {1 42 144}
	// 3
}

func Example3() {
	var x IntSet
	x.Add(1)
	x.Add(1)
	x.Add(100)
	fmt.Println(x.String()) // "{1 100}"
	fmt.Println(x.Len())    // "2"

	x.Clear()
	fmt.Println(x.String()) // "{}"
	fmt.Println(x.Len())    // 0

	// Output:
	// {1 100}
	// 2
	// {}
	// 0
}

func Example4() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y := x.Copy()
	fmt.Println(y.String()) // "{1 9 144}"
	y.Add(11)
	fmt.Println(x.String()) // "{1 9 144}"
	fmt.Println(y.String()) // "{1 9 11 144}"

	// Output:
	// {1 9 144}
	// {1 9 144}
	// {1 9 144}
	// {1 9 11 144}
}

func Example5() {
	var x, y IntSet
	x.AddAll(1, 9, 144)
	fmt.Println(x.String()) // "{1 9 144}"

	y.AddAll(3, 4, 7)
	y.UnionWith(&x)
	fmt.Println(y.String()) // "{1 3 4 7 9 144}"

	// Output:
	// {1 9 144}
	// {1 3 4 7 9 144}
}

func Example6() {
	var x, y IntSet
	x.AddAll(1, 9, 144)
	y.AddAll(3, 9, 1)
	y.IntersectWith(&x)
	fmt.Println(y.String)

	// Output
	// {1 9}
}

func Example7() {
	var x, y IntSet
	x.AddAll(1, 9, 144)
	y.AddAll(3, 9, 4)
	x.DifferenceWith(&y)
	fmt.Println(x.String)

	// Output
	// {1 144}
}

func Example8() {
	var x, y IntSet
	x.AddAll(1, 9, 144)
	y.AddAll(3, 9, 4)
	x.SymmetricDifferenceWith(&y)
	fmt.Println(x.String)

	// Output
	// {1 3 4 144}
}

func Example9() {
	var x IntSet
	x.AddAll(1, 9, 144)
	elems := x.Elems()
	areEqual := len(elems) == 3 && elems[0] == 1 &&
		elems[1] == 9 && elems[2] == 144
	fmt.Println(areEqual)

	// Output
	// true
}
