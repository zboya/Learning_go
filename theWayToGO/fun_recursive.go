package main

func main(){
	for i:=10;i>=1;i-- {
		print(recursive(i))
	}
}

func recursive(n int) (res int) {
	if n==10 {
		res=10
		return 
	}else{
		res=recursive(n+1)-1
	}
	return 
}
