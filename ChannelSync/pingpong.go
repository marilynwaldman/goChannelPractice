package ChannelSync

import (
	fmt "fmt"

)

func PingPong() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, pongs)
	pong(pings, pongs)
	fmt.Println(<-pongs)

}

func ping(pings chan<- string,pongs <-chan string ) {
	pings <- "ping"
	//msg := <-pongs
	//fmt.Println(msg)
}

func pong(pings <-chan string, pongs chan<- string) {

	msg := <-pings
	fmt.Println(msg)
	pongs <- "pongs"


}