package main

import (
    "fmt"
    
    "exercises/concurrency/pipeline"
)

func main() {
    // Set up the pipeline.
    c := pipeline.Gen(2, 3)
    out := pipeline.Sq(c)

    // Consume the output.
    fmt.Println(<-out) // 4
    fmt.Println(<-out) // 9
}