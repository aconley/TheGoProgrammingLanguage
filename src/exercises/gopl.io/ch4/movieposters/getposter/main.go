package main

import (
	"fmt"
	"log"
	"os"

	"exercises/gopl.io/ch4/movieposters"
)

// getposter returns the urls for movies matching the provided names
func main() {
    if len(os.Args) < 2 {
        log.Fatal(fmt.Errorf("Usage: getposter title1 [title2 ...]"))
    }
    for _, name := range(os.Args[1:]) {
        result, err := movieposters.GetPosterURLs(name)
        if err != nil {
		  log.Fatal(err)
        }
        fmt.Printf("For the movie '%s' found %d posters:\n",
            name, len(result))
        for _, url := range(result) {
            fmt.Printf("\t%s\n", url)
        }
	}
}