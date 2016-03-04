package main

import (
    "fmt"
    
    "exercises/concurrency/pipeline"
)

func main() {
    // Set up a done channel that's shared by the whole pipeline,
    // and close that channel when this pipeline exits, as a signal
    // for all the goroutines we started to exit.
    done := make(chan struct{})
    defer close(done)

    in := pipeline.Gen(2, 3)

    // Distribute the sq work across two goroutines that both read from in.
    c1 := pipeline.SqWithDone(done, in)
    c2 := pipeline.SqWithDone(done, in)

    // Consume the first value from output.
    //  Note this will only print one value!
    out := pipeline.MergeWithDone(done, c1, c2)
    fmt.Println(<-out) // 4 or 9

    // done will be closed by the deferred call.
}