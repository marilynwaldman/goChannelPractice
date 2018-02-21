package ChannelSync

import "fmt"

func worker4(c chan int){
	c <- 1

}

func Sync(){
	c := make(chan int)
	go worker4(c)
	<-c
	fmt.Println("done")
}
