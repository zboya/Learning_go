package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length,width float32	
} 

func (r Rectangle) Area() float32 {
	return r.length * r.width
}


func main(){
	sq1:=new(Square)
	sq1.side=5
	areaint:=sq1
	fmt.Printf("the square has area :%f\n",areaint.Area())
	fmt.Printf("the square has area :%f\n",sq1.Area())

	r:=Rectangle{3,5}
	fmt.Printf("the rectangle has area :%f\n",r.Area())
}


