package main

import(
	"fmt"
	"strings"

)

func main(){
	//srings.HasPrefix()
	var str string = "this is: an example of a string"
	fmt.Printf("T/F? Dose the string \"%s\" have prefix %s? ",str,"TH")
	fmt.Printf("%t\n",strings.HasPrefix(str,"TH"))
	//strings.Contains()
	str2:="sheepbao is learning the golang language"
	fmt.Printf("\"%s\" contains \"the\" ? ",str2)
	fmt.Printf("%t\n",strings.Contains(str2,"the"))
	//strings.Index(s,str string) int
	//strings.LastIndex() int 
	//strings.replace(str,old,new,n) string
	//strings.Count( ) int

	fmt.Printf("Number of i in %s is : %d \n",str,strings.Count(str,"i"))
	fmt.Printf("replace an to many %s\n",strings.Replace(str,"an","many",-1))
	//strings.Fields()
	fmt.Printf("%v\n",strings.Fields(str))
	fmt.Printf("%v\n",strings.Split(str,":"))

	

}
