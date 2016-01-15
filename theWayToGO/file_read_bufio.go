package main

import (
	"fmt"
	"bufio"
	"io"
	"os"
)

func main(){
	inputFile,inputError := os.Open("input.dat")
	if inputError!=nil {
		fmt.Printf("file Open failed\n")
		return 
	}
	defer inputFile.Close()

	inputReader:=bufio.NewReader(inputFile)
	for {
		inputString,readerError := inputReader.ReadString('\n')
		if readerError == io.EOF {
			println("EOF")
			//return
			break
		}
		fmt.Printf("the input was : %s ",inputString)
	}
}
