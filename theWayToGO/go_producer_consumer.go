package main
import(
"runtime"
)
func producer(ch chan int) {
	for i:=0;i<10;i++ {
		ch <- (i*10)
	}
	close(ch)
}

func consumer(ch chan int , done chan bool) {
	for n:=range ch{
		println("i am consumer one")
		println(n)
	}
	println("is done")
	done <- true
}
func consumer2(ch chan int , done chan bool) {
	//rec 0 forever
	for {
		println("i am consumer two")
		println(<-ch)
	}
	println("is done")
	done <- true
}

func main() {
	println(runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU())
	ch:=make(chan int)
	done1:=make(chan bool)
	done2:=make(chan bool)
	go producer(ch)
	go consumer(ch,done1)
	go consumer2(ch,done2)
	<-done1
	<-done2
}
