// Wallclock is a TCP client that echoes several
//  instances of clock2 to a table.
package main

import (
  "bufio"
  "fmt"
  "strings"
  "log"
  "net"
  "os"
)

func getTimes(readers []bufio.Reader) ([]string, error) {
  n := len(readers)
  results := make([]string, n, n)
  for idx, rdr := range(readers) {
    time, err := rdr.ReadString('\n')
    if err != nil {
      return nil, err
    }
    results[idx] = time[:len(time)-1]
  }
  return results, nil
}

func handleConns(names []string, cs []net.Conn) {
  // The tricky part is that we want to do this
  //  without channels
  n := len(cs)
  fmt.Printf("Number of conns: %d\n", n)
  if n == 0 {
    return
  }

  // Convert the connections into readers
  readers := make([]bufio.Reader, n, n)
  for i, c := range(cs) {
    readers[i] = *bufio.NewReader(c)
  }

  // Main loop
  for {
    times, err := getTimes(readers)
    if err != nil {
      return
    }
    for idx, time := range(times) {
      fmt.Printf("\t%s:%s", names[idx], time)
    }
    fmt.Printf("\n")
  }
}

func main() {
  n := len(os.Args) - 1

  if n == 0 {
    fmt.Print("Usage: wallclock name=port [name=port ...]")
    return
  }

  names := make([]string, n, n)
  conns := make([]net.Conn, n, n)

  // Initialize
  for idx, val := range(os.Args[1:]) {
    arg := strings.SplitN(val, "=", 2)
    if len(arg) != 2 {
      log.Fatal(fmt.Sprintf("Couldn't split %s", arg))
    }
    names[idx] = arg[0]
    conn, err := net.Dial("tcp", arg[1])
    if err != nil {
      log.Fatal(err)
    }
    conns[idx] = conn
    defer conn.Close()
  }

  handleConns(names, conns)
}