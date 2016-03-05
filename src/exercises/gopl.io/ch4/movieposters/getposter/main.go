package main

import (
	"fmt"
	"log"
	"os"

	"exercises/gopl.io/ch4/movieposters"
)

// getposter returns the urls for movies searching on the provided
//  titles
func main() {
  if len(os.Args) < 2 {
      log.Fatal(fmt.Errorf("Usage: getposter title1 [title2 ...]"))
  }
  for _, name := range(os.Args[1:]) {
    urls, err := movieposters.GetPosterURLs(name)
    if err != nil {
      log.Fatal(err)
    }
    
    outnames, err := movieposters.WritePosters(name, urls)
    if err != nil {
      log.Fatal(err)
    }
    
    fmt.Printf("For the movie '%s' found %d posters\n",
        name, len(outnames))
    if len(outnames) > 0 {
      fmt.Printf(" Written to:\n")
      for _, outname := range(outnames) {
        fmt.Printf("\t%s\n", outname)
      }
    }
  }
}