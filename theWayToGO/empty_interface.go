package main

import "fmt"

var i=5
var str="ABC"

type Person struct {
	name string
	age int
}

type Any interface {}

func main(){
	var val Any
	val=5
	fmt.Printf("val has the value %v\n",val)
	val = str
	fmt.Printf("val has the value %v\n",val)
	p1:=new(Person)
	p1.name="Rob Pike"
	p1.age=25
	val = p1
	fmt.Printf("val has the value %v\n",val)
	switch t:=val.(type) {
		case int:
			fmt.Printf("type int %T\n",t)
		case string:
			fmt.Printf("type string %T\n",t)
		case bool:
			fmt.Printf("type bool %T\n",t)
		case *Person:
			fmt.Printf("type pointer person %T\n",t)
		default:
			fmt.Printf("Unexpected type %T\n",t)
		}

}
