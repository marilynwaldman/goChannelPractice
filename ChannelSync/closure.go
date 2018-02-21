package ChannelSync

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}

}

func MyClosure(){
	pos := adder()
	for i := 1; i <= 10; i++{
		fmt.Println(pos(i))
	}

}