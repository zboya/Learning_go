package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("in main")
	go longWait(5)
	go shortWait(2)
	println("sleep in main()")
	time.Sleep(4 *1e9)
	println("end of main()")
}

func longWait(n int) {
	fmt.Printf("Begining longWait()\n")
	n=n*1e9
	time.Sleep(10)
	fmt.Printf("End of longWait()\n")
}

func shortWait(n int) {
	fmt.Printf("Beginning shortWait()\n")
	n=n*1e9
	time.Sleep(2)
	fmt.Printf("End of shortWait()\n")
}
