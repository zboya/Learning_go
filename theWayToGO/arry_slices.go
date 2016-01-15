package main
import "fmt"
func main() {
	var arry1 [6]int
	var slice1 []int = arry1[2:5]

	for i := 0; i < len(arry1); i++ {
		arry1[i]=i
	}

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice1 at %d is %d\n",i,slice1[i])
	}

	fmt.Printf("the len of arry1 is %d\n",len(arry1))
	fmt.Printf("the len of slice1 is %d\n",len(slice1))
	fmt.Printf("the cap of slice1 is %d\n",cap(slice1))

	slice1=slice1[0:4]

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("slice1 at %d is %d\n",i,slice1[i])
	}

	fmt.Printf("the len of slice1 is %d\n",len(slice1))
	fmt.Printf("the cap of slice1 is %d\n",cap(slice1))
	println(sum(slice1[:]))
}


func sum(s []int) int {
	res:=0
	for i := 0; i < len(s); i++ {
		res+=s[i]
	}
	return res
}