package ChannelSync

import "fmt"

func ClosingChannels(){

	jobs := make(chan int, 5)
	done := make(chan bool)

	go func(){
		for {
			j, more := <- jobs
			if(more){
				fmt.Println("job recieved ", j)
			} else {
				fmt.Println("No more jobs")
				done <- true
				return
			}
		}

	}()

	for j:= 1; j <4 ; j++{
		jobs <- j
		fmt.Println("sent job j ", j)

	}
	fmt.Println("Sent all jobs s")
	close(jobs)
	<- done

}
