package main

import (
    "fmt"
    
    "exercises/concurrency/pipeline"
)

func main() {
    in := pipeline.Gen(2, 3)

    // Distribute the sq work across two goroutines that both read from in.
    c1 := pipeline.Sq(in)
    c2 := pipeline.Sq(in)

    // Consume the merged output from c1 and c2.
    for n := range pipeline.Merge(c1, c2) {
        fmt.Println(n) // 4 then 9, or 9 then 4
    }
}