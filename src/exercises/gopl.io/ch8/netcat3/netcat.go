// Netcat is a read-only tcp client
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

	defer connTCP.Close()

	done := make(chan struct{})
	go func() {
		if _, err := io.Copy(os.Stdout, conn); err != nil {
			done <- struct{}{}
			log.Fatal(err)
		}
		log.Println("done")
		done <- struct{}{}
	}()

	connTCP.CloseWrite()
	<-done // wait for background connection to finish
}
