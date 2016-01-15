package main

import "fmt"

func main(){
	fmt.Printf("%d is even: is %t\n",16,even(16))
	fmt.Printf("%d is odd: is %t\n",17,odd(17))
	fmt.Printf("%d is odd: is %t\n",18,odd(18))
}

func even(n int) (bool) {
	print("even ")
	if n==0 {
		return true
	}
	return odd(RevSign(n)-1)
}

func odd(nr int) (bool) {
	print("odd ")
	if nr==0 {
		return false
	}
	return even(RevSign(nr)-1)
}

func RevSign(nr int)(int) {
	if nr<0 {
		return -nr
	}
	return nr
}
