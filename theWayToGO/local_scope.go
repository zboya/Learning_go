package main

var a = "G"

func main(){
	n()
	m()
	n()
}

func n(){print(a)}

func m(){
	//attention
	a:="O"
	print(a)
}
