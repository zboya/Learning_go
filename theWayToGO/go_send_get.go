package main

func main() {
	println("main")
	ch:=make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	println("senddata")
	ch <- "sheep"
	ch <- "bao"
	close(ch)
}

func getData(ch chan string) {
	println("getdata")
	for {
		input,open:= <-ch
		if !open {
			break
		}
		println(input)
	}
}
