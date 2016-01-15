package main

func main() {
	for i := 0; i < 30; i++ {
		println(i,"!=",factorial(i))
	}
}

func factorial(n int) (res int) {
	if n==0 {
		res=1
		return
	}else{
		res=n*factorial(n-1)
	}
	return

}