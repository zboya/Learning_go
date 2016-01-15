package main

type struct1 struct {
	i1 int
	f1 float32
	str string
}

func main(){
	ms:=new(struct1)
	ms.i1=10
	ms.f1=15.5
	ms.str="chris"

	println(ms.i1)
	println(ms.f1)
	println(ms.str)
	print(ms)
}
