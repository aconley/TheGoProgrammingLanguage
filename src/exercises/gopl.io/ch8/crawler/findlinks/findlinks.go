// findlinks crawls web links starting with the command-line arguments.
//
// This version uses a buffered channel as a counting semaphore
// to limit the number of concurrent calls to links.Extract.
//
// This version is modified from the one in the book (p 241)
//  to allow a maximum depth, as per exercise 8.6
package main

import (
	"flag"
	"fmt"
	"log"

	"exercises/gopl.io/ch8/crawler"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

type linksPlusDepth struct {
	links []string
	depth int
}

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	list, err := crawler.Extract(url)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {

	maxDepthFlag := flag.Int("depth", 1000, "maximum depth (def: 1000)")
	flag.Parse()
	maxDepth := *maxDepthFlag

	worklist := make(chan linksPlusDepth)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() {
		worklist <- linksPlusDepth{flag.Args(), 0}
	}()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		item := <-worklist
		if item.depth < maxDepth {
			for _, link := range(item.links) {
				if !seen[link] {
					seen[link] = true
					n++
					go func(link string, depth int) {
						worklist <- linksPlusDepth{crawl(link), depth + 1}
					}(link, item.depth)
				}
			}
		} else {
			// We mark them as seen... but don't send them
			for _, link := range(item.links) {
				if !seen[link] {
					seen[link] = true
				}
			}
		}
	}
}
