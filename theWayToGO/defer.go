package main

//import 

func main(){
	f1()
	a()
	f()
}

func f1(){
	print("In f1 at the top\n")
	defer f2()
	print("In f1 at the bottom\n")

}

func f2(){
	print("f2:Deferred until the end of the calling function!\n")
}

func a(){
	i:=0
	defer print(i,"\n")
	i++
	return
}

func f(){
	for i:=0;i<5;i++ {
		defer print(i)
	}
	print("\n")
}
