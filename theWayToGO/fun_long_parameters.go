package main
import "fmt"
func main() {
	x:=min(1,4,5,7,9,3)
	print(x,"\n")
	arr:=[]int{1,4,5,8}
	x=min(arr...)
	print(x,"\n")
}

func min(a ...int) (int) {
	//fmt.Printf("%T",a)
	if len(a)==0 {
		return 0
	}
	m:=a[0]
	for _,v := range a {
		if v< m {
			m=v
		}
	}
	return m
}