//how it work?
package main

func generate(ch chan int) {
	for i:=2;i<1000;i++ {
		ch <- i
	}
}

func filter(in,out chan int,prime int,done chan bool) {
	for i:=range in {
		if i%prime !=0 {
			out <- i
		}
	}
	done <- true
}
func main() {
	ch:=make(chan int)
	go generate(ch)
	for i:= range ch {
		prime:= i
		print(prime," ")
		ch1:=make(chan int)
		done:=make(chan bool)
		go filter(i,ch1,prime,done)
		i=ch1
		<-done
	}
}
