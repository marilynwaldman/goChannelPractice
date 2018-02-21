package ChannelSync

import "fmt"

type vertex struct {
	x,y int
}

func (v vertex) adder() int {
	return v.x + v.y

}

func MyMethod(){
	v := vertex{3,4}
	fmt.Println(v.adder())
}