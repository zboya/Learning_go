package main

import "fmt"

type TwoInt struct{
	a int
	b int
}

func main(){
	t1:=&TwoInt{12,10}
	fmt.Printf("the sum is %d\n",t1.AddThem())
	fmt.Printf("Add them to the param : %d\n",t1.AddToParam(20))

	t2:=TwoInt{3,4}
	fmt.Printf("the sum is %d \n",t2.AddThem())

}

func (tn *TwoInt) AddThem() int {
	return tn.a+tn.b
}

func (tn *TwoInt) AddToParam (param int) int {
	return tn.a+tn.b+param

}
