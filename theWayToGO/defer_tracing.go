package main

func trace(s string ) {print("entering: ",s)}
func untrace(s string) {print("leaving:",s)}

func a(){
	trace("a()")
	defer untrace("a()")
	print("in a()")
}

func b(){
	trace("b()")
	defer untrace("b()")
	print("in b()")
	a()
}

func main(){
	b()
}

