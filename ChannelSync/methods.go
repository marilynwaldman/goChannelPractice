package ChannelSync

import (
	"fmt"
)

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r *rect) perimeter() int {
	return 2 * r.height + 2 * r.width
}

func Method() {
	r := rect{width: 10, height: 5}
	fmt.Println("area ", r.area())
	fmt.Println("area ", r.perimeter())

	rd := &r

	fmt.Println("area ", rd.area())
	fmt.Println("area ", rd.perimeter())


}








