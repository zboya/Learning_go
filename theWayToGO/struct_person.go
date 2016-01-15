package main

import "strings"

type person struct {
	name string
}

func main(){
	//1
	var p1 person
	p1.name="zhangsan"
	upper(&p1)
	println(p1.name)
	//2
	p2:=new(person)
	p2.name="lisi"
	println((*p2).name)
	//3
	p3:=&person{"wangmazi"}
	println(p3.name)

}

func upper(p *person) {
	p.name=strings.ToUpper(p.name)
}
