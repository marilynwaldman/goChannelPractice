package ChannelSync

import (
	"time"
	"fmt"
)

type Ball struct {
	hits int
}

func PingPong2(){
	table := make(chan *Ball)
	go player("ping", table)
	go player("pong", table)

	table <- new(Ball)

	time.Sleep(1 * time.Second)

	<-table
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
