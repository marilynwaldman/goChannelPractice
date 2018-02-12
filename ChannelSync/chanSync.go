package ChannelSync

import (
	"fmt"
	"time"
)

func ChanSync(){

	done := make(chan bool, 1)
	go worker(done)
	<- done

}

func worker(done chan bool){
	fmt.Println("working ....")
	time.Sleep(time.Second)
	done <- true
	fmt.Println("Done")


}
