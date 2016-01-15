package main

func main(){
	for i:=0;i<5;i++ {
		print(i)
	}
	print("\n")
	for i,j:=0,10;i<j;i,j=i+1,j-1 {
		print(i,j,"\n")
	}
}
