package main

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	ch <- "sheep"
	ch <- "bao"
}

func getData (ch chan string) {
	var input string
	for {
		input = <- ch
		fmt.Printf("%s\n",input)
	}
}

func main() {
	ch:=make (chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1*1e9)
}
