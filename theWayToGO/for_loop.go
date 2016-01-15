package main

func main(){
	for i:=0;i<15;i++ {print(i)}
	print("\n")
	n:=0
loop:
	if n<15 {
		print(n,"\n")
		n++
		goto loop
	} 
//	print(n)
}
