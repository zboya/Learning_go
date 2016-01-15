package main

import "fmt"

func main() {
	p2:=Add2()
	fmt.Printf("call add2 for gives %v\n:",p2(3))
	Twoadder:=Adder(2)
	fmt.Printf("the result is %v\n",Twoadder(3))
}

func Add2() func(b int) int {
	// print("add2",b)
	return func(b int) int{
		println("add2in",b)
		return b+2
	}
}

func Adder(a int) func(b int) int {
	print("adder",a)
	return func(b int) int {
		println("adderin",b)
		return a+b
	}
}
