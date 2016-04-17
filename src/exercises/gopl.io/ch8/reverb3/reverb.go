// Reverb3 is a TCP server that simulates an echo.
//  This version is modified as per exercise 8.4 and 8.8
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

	inputChan := make(chan string)

	// Puts text on inputChan
	go func(connTCP *net.TCPConn) {
		input := bufio.NewScanner(c)
		for input.Scan() {
			inputChan <- input.Text()
		}
	}(connTCP)

  var wg sync.WaitGroup // Number of open echos per connection

	timeOutTicker := time.NewTicker(10 * time.Second)

	select {
		case text := <- inputChan:
			log.Println("Starting new echo")
			wg.Add(1)
			go func(txt string, delay time.Duration) {
      	defer wg.Done()
      	fmt.Fprintln(connTCP, "\t", strings.ToUpper(txt))
	    	time.Sleep(delay)
	    	fmt.Fprintln(connTCP, "\t", txt)
	    	time.Sleep(delay)
	    	fmt.Fprintln(connTCP, "\t", strings.ToLower(txt))
    	}(text, 1 * time.Second)
		case <-timeOutTicker.C:
			log.Println("Heard nothing on channel; closing")
			timeOutTicker.Stop()
			connTCP.Close()
			return
	}
	timeOutTicker.Stop()

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
