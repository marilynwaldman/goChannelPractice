package ChannelSync

import (
	"fmt"
	"time"
)

func pinger(c chan string){
	fmt.Println("pinger")
	for i:=0;;i++{
		c <- "hi"

	}
}

func printer(c chan string){
	for {
		s := <- c
		fmt.Println(s)
		time.Sleep(time.Second)
	}
}

func Pinger(){
	var c chan string = make(chan string)
	fmt.Println("here")

	go pinger(c)
	go printer(c)
	time.Sleep(10 * time.Second)
}
