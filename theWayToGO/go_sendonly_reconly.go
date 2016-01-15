package main

func main() {
	var c=make(chan int)
	go source(c)
	go sink(c)
}

func source(ch chan<- int) {
	for {ch <- 1}
}

func sink(ch <-chan int) {
	for { <-ch }
}

