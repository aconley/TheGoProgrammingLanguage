// Reverb3 is a TCP server that simulates an echo.
//  This version is modified as per exercise 8.4
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
  "sync"
	"time"
)

func handleConn(c net.Conn) {
  var connTCP = c.(*net.TCPConn)

  var wg sync.WaitGroup // Number of open echos per connection
	input := bufio.NewScanner(c)
	for input.Scan() {
    log.Println("Starting new echo")
    wg.Add(1)

    go func(txt string, delay time.Duration) {
      defer wg.Done()
      fmt.Fprintln(connTCP, "\t", strings.ToUpper(txt))
	    time.Sleep(delay)
	    fmt.Fprintln(connTCP, "\t", txt)
	    time.Sleep(delay)
	    fmt.Fprintln(connTCP, "\t", strings.ToLower(txt))
    }(input.Text(), 1 * time.Second)
	}

  // Closer
  go func() {
    wg.Wait()
    connTCP.CloseWrite()
  }()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
