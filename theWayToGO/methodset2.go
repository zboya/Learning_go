package main

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l=append(*l,val)
}

type Appender interface {
	Append(int)
}

func CountInto (a Appender,start,end int){
	for i:=start;i<=end;i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEough(l Lener) bool {
	return l.Len()*10 > 42
}

func main(){
	var lst List
	if LongEough(lst) {
		fmt.Printf("-lst is long enough\n")
	}

	plst:=new(List)
	CountInto(plst,1,10)
	if LongEough(plst) {
		fmt.Printf("-plst is long enough\n")
	}
}
