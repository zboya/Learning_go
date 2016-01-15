package main

func main() {
	m1:=make(map[string]int)

	m1["one"]=1
	m1["two"]=2

	println(m1["one"])
	println(m1["two"])
}