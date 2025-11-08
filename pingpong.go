package main

import (
	"fmt"
	"time"
)

type Ball struct{ hits int }

func main() {
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball)
	time.Sleep(1 * time.Second)
	<-table

	panic("show me the stacks") // GOTRACEBACK=all go run pingpong.go

	// OR:
	//
	// buf := make([]byte, 1024*1024)
	// stackSize := runtime.Stack(buf, true)
	// fmt.Printf("\n%s\n", buf[:stackSize])
}

func player(name string, table chan *Ball) {
	for {
		ball := <-table
		ball.hits++
		fmt.Println(name, ball.hits)
		time.Sleep(100 * time.Millisecond)
		table <- ball
	}
}
