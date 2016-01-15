package main

import "time"


func main(){
	tick:=time.Tick(1e8) //0.1s
	boom:=time.After(5e8) // 0.5s

	for {
		select {
			case <-tick:
				println("tick.")
			case <-boom:
				println("boom.")
				return
			default:
				println(" .")
				time.Sleep(5e7) //0.01s
		}
	}
}
