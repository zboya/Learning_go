package main

func main() {
	s1_from:=[]int{1,2,3}
	s1_to:=make([]int,10)

	n:=copy(s1_to,s1_from)
	println(s1_to)
	println(n)

	s13:=[]int {1,3,4}
	s13=append(s13,4,5,6)
	println(s13)

	for i := 0; i < 6; i++ {
		println(s13[i])
	}
}