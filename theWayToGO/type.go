package main

import "fmt"

type TZ int

func main(){
	var a,b TZ =3,4
	c:=a+b
	fmt.Printf("c=%d\n",c)
	//print(c)
	
	var c1 byte = 65
	var c2 byte = '\x41'
	var ch int ='\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'

	fmt.Printf("%d-%d-%d-%d-%d\n",c1,c2,ch,ch2,ch3)
	fmt.Printf("%c-%c-%c-%c-%c\n",c1,c2,ch,ch2,ch3)
	fmt.Printf("%X-%X-%X-%X-%X\n",c1,c2,ch,ch2,ch3)
	fmt.Printf("%U-%U-%U-%U-%U\n",c1,c2,ch,ch2,ch3)
}
