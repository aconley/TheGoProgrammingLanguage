package main

import (
	"fmt"
	"time"
)

type ball struct{ hits int }

func main() {
	table := make(chan *ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(ball)
	time.Sleep(1 * time.Second)
	<-table
}

func player(name string, table chan *ball) {
	for {
		b := <-table
		b.hits++
		fmt.Println(name, b.hits)
		time.Sleep(100 * time.Millisecond)
		table <- b
	}
}
