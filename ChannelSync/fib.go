package ChannelSync

import (
	"fmt"
	"time"
)

func Fib(){
	go spinner(100* time.Millisecond)
	const n = 45;
	nums := make(map[int] int)
	nums[0] = 0
	nums[1] = 1
	fibN := fib(n)
	fmt.Println(fibN);

}

func spinner(delay time.Duration){
	for {
		for _,r := range `-\|/`{
			time.Sleep(delay)
			fmt.Printf("\r%c", r)

		}
	}
}

func fib(n int) int {

	if( n < 2){
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}

}
