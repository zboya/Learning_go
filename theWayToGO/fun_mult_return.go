package main

func main() {
	a,b:=f1(2)
	print(a,b)
	c,d:=f2(2)
	print(c,d)
}

func f1(input int) (int , int) {
	return input*2,input*3
	
}
func f2(input int) (x2 int,x3 int) {
	x2=input*2
	x3=input*3
	return
}