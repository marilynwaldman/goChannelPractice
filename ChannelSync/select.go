package ChannelSync

import (
	"fmt"
	"time"
)

func SelectChan(){

	c1 := make(chan string, 1)
	c2 := make(chan string, 1)

	go func(){
		time.Sleep(1.0*time.Millisecond)
		c1 <- "one"
	}()

	go func(){
		time.Sleep(2.0*time.Millisecond)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++{
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}
