package main

import (
	"fmt"
	"os"
)
//variables go
	//define
	var a int
	var b bool
	var str string

	//also declare

	var(
		a2 int
		b2 bool
		str2 string
	)

	// a=15
	// b=false

	//declaration and assignment can of course be combined
	var a3 int = 15
	var i = 5
	var b3 bool = false 
	var str3 string ="go say hello to the world!"

	//also right

	var a4=15
	var b4=false
	var str4="go say hello to the world!"
	//
	// short form :=  only in function body
	// a5:=50
	// b5:=false
	//one , two , three := 1,2,3
	//	
func main(){
	var goos string = os.Getenv("GOOS")
	fmt.Printf("the operating system is %s\n",goos)
	path:=os.Getenv("PATH")
	fmt.Printf("Path is %s\n",path)
}
