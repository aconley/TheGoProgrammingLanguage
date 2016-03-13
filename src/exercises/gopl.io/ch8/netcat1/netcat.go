// Netcat1 is a read-only tcp client
package main

import (
  "io"
  "flag"
  "fmt"
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
  defer conn.Close()
  mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
  if _, err := io.Copy(dst, src); err != nil {
    log.Fatal(err)
  }
}