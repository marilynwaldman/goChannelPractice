package ChannelSync

import (
	"time"
	"fmt"
)

func RateLimit(){

	requests := make(chan int, 50)
	for i := 1; i <=5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick(2 * time.Second)

	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

}