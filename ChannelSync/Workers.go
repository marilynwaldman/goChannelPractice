package ChannelSync

import (
	"fmt"
	"time"
)

func poolWorker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func  WorkerPool(){
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w:= 1; w <= 3; w++ {
		go poolWorker(w, jobs, results)
	}

	for j:= 1; j <= 5; j++{
		jobs <- j
	}

	close(jobs)

	for a := 1; a <= 5; a++ {
		<-results
	}
}
