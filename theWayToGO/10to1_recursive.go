package main

func main(){
	recursive(10)
}

func recursive(n int) int {
	println(n)
	if n>1 {
		return recursive(n-1)
	}
	return 1 
}