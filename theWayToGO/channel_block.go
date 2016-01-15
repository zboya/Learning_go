package main
import "time"
func main() {
	ch1:=make(chan int)
	go pump(ch1)
	go suck(ch1)
	time.Sleep(1e9)
	//println(<-ch1)
}


func pump(ch chan int) {
	for i:=0;i<20000;i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for{
		println(<-ch)
	}
}
