package main

import (
        "fmt"
        "os"
        "strconv"

        "exercises/gopl.io/ch2/tempconv"
)

func main() {
        for _, arg := range os.Args[1:] {
                t, err := strconv.ParseFloat(arg, 64)
                if err != nil {
                        fmt.Fprintf(os.Stderr, "cf: %v\n", err)
                        os.Exit(1)
                }
                f := tempconv.Fahrenheit(t)
                c := tempconv.Celsius(t)
                k := tempconv.Kelvin(t)
                fmt.Printf("%s = %s, %s = %s, %s = %s = %s\n",
                        f, tempconv.FToC(f), c, tempconv.CToF(c),
                        k, tempconv.KToC(k), tempconv.KToF(k))
        }
}


