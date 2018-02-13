package ChannelSync

import (
	"time"
	"fmt"
)

func TimeOut(){

	c1 := make(chan string, 1)
	go func() {
		time.Sleep(8 * time.Millisecond)
		c1 <- "one"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Millisecond):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Millisecond)
		c2 <- "two"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 2")
	}
}
