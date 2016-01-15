package main

func main(){
	var num int =7
	switch {
		case num < 0:
			print("NUM is negative")
		case num > 0 && num < 10 :
			print("NUM is between 0 to 10")
		default:
			print("default")
	}
}
