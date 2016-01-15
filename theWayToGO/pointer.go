package main

import "fmt"

func main(){
	var i1=5
	fmt.Printf("i1's address is : %p \n",&i1)
	var intp *int
	intp=&i1
	fmt.Printf("%d pointer is %p ",*intp,intp)
}
