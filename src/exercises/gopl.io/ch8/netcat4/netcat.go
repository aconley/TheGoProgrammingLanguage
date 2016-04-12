// Netcat is a read-only tcp client
// This version is modified as per exercise 8.3

// I decided not to have this keep running after
//  main is killed.  So the way to test what this does
//  when stdin is closed is to pipe in input:
//   netcat4 < input_file
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	var fport = flag.Int("port", 8000, "port")
	var fhost = flag.String("host", "localhost", "host")

	flag.Parse()
	connString := fmt.Sprintf("%s:%d", *fhost, *fport)

	conn, err := net.Dial("tcp", connString)
	if err != nil {
		log.Fatal(err)
	}
	var connTCP = conn.(*net.TCPConn)

	done := make(chan struct{})
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			log.Fatal(err)
		}
		done <- struct{}{}
	}()

  mustCopy(conn, os.Stdin)
	connTCP.CloseWrite()
	<-done // wait for background connection to finish
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}