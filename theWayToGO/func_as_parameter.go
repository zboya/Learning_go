package main


func main() {
	callback(2,add)
}

func add(a,b int) {
	println("a+b=",a+b)
	
}

func callback(y int,f func(int,int)) {
	f(y,2)
}