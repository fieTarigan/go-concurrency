package main

import (
	"fmt"
	"strconv"
	"time"
)

func newpinger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "ping-" + strconv.Itoa(i)
	}
}

func newponger(c chan<- string) {
	for i := 0; ; i++ {
		c <- "pong-" + strconv.Itoa(i)
	}
}

func newprinter(c <-chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string, 4)

	go newpinger(c)
	go newponger(c)
	go newprinter(c)

	var input string
	fmt.Scanln(&input)
}
