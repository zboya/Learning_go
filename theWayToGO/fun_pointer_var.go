package main

func main() {
	n:=0 ;r:=&n
	multiply(2,3,r)
	print(*r)
}

func multiply(a,b int, reply *int) {
	*reply=a*b
}