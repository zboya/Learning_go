package main

import "fmt"
func main() {
	str:="the way to go\n"
	print(str)
	for k,v:=range str {
		fmt.Printf("%d, %c \n",k,v)
	}
}