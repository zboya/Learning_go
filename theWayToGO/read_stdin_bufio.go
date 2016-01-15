package main

import (
	"fmt"
	"bufio"
	"os"
)

var inputReader *bufio.Reader
var input string
var err error

func main(){
	inputReader = bufio.NewReader(os.Stdin)
	fmt.Println("please enter some input:")
	input,err=inputReader.ReadString('\n') //end of string
	if err==nil {
		fmt.Printf("The input was :%s\n",input)
	}
}
