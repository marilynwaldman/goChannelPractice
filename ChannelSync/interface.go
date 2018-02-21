package ChannelSync

import (
	"fmt"
	"math"
)

type geometry interface {
	area()
	perim()
}

type rectI struct{
	width float32
	height float32
}

type circleI struct{
	radius float32
}

func (r rectI) areaI() float32{
	return r.width * r.height
}

func (r circleI) areaI() float32{
	return r.radius * 2 * math.Pi
}
func Interface() {
	r := rectI{width: 10, height: 5}
	c := circleI{radius: 2}
	fmt.Println("Rectangle area ",r.areaI())
	fmt.Println("Circle ", c.areaI())
}


