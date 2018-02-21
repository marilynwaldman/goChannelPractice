package ChannelSync

import (
	"time"
	"fmt"
)

func RateLimit2(){
	requests := make(chan int,5)
	for i := 1; i<= 5; i++ {
		requests <- i
	}

	close(requests)

	limiter := time.Tick( 2 * time.Millisecond)

	for req := range requests{
		<- limiter
		fmt.Println(req)

	}

	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	go func() {
		for t := range time.Tick(2 * time.Second){
			burstyLimiter <- t
		}
	}()

	burstyRequests := make(chan int, 12)

	for i := 1; i <=12; i++ {
		burstyRequests <- i
	}

	for req := range burstyRequests{
		<- burstyLimiter
		fmt.Println("request ", req, time.Now())

	}

	close(burstyRequests)


}
