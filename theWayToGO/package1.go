package main

import "./pack1"

func main(){
	var test1 string
	test1=pack1.ReturnStr()
	println(test1)
	println("int from pack1",pack1.Pack1Int)
}
