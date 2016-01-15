package main

func main(){
	print(Abs(2))
	print("\n")
	print(isGreater(5,6))
}


func Abs(x int) int{
	if x<0 {
		return -x
	}else{
		return x
	}
}

func isGreater(x,y int) bool{
	if x>y {
		return true
	}
	return false
}
