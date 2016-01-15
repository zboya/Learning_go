package main

import "strings"

func main(){
	for i:=1;i<=25;i++ {
		for j:=0;j<i;j++{
			print("G")
		}
		print("\n")
	}

	print("---------one for loop---------\n")
	for i:=1;i<=25;i++ {
		print(strings.Repeat("G",i),"\n")
	}
	print("---------one for loop---------\n")
	var str string="G"
	for i:=1;i<=25;i++ {
		print(str,"\n")
		str=str+"G"
	}
}
