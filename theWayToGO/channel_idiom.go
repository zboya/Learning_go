package main

//import "time"

func main() {
	done:=make(chan bool)
	stream:=pump()
	go suck(stream,done)
	//time.Sleep(1e9)
	<-done
}

func pump() (chan int) {
	ch:=make(chan int)
	go func(){
		for i:=0;i<20000;i++ {
			ch <- i
		}
	}()
	return ch
}

func suck(ch chan int,done  chan bool) {
	for {
		println(<-ch)
	}
	done <- true
}
