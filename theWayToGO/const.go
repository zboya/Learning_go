package main

import "fmt"

func main() {
    const Pi=3.14159
	const Monday,Tuesday,Wednesday,Thursday,Friday,Saturday=1,2,3,4,5,6
	const (
		One,Two,Three=1,2,3
		Four,Five,Six=4,5,6
	)
	type Color int

	const (
		RED Color = iota //0
		ORANGE //1
		YELLOW
		GREEN
		BLUE
		INDIGO
		VIOLET //6
	)
	//fmt.Printf(Monday)
	fmt.Printf("hello, world")
	//test()
}

//func test(){
//	fmt.Printf("this is a test\n")
//}

