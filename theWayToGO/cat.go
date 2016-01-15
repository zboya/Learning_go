package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"

)

func cat(r *bufio.Reader) {
	for {
		buf,err:=r.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Fprintf(os.Stdout,"%s",buf)

	}
	return 
}

func main() {
	flag.Parse()
	if flag.NArg()==0 {
		cat(bufio.NewReader(os.Stdin))
	}
	for i:=0;i<flag.NArg();i++ {
		f,err := os.Open(flag.Arg(i))
		if err!=nil {
			println("error")
			continue
		}
		cat(bufio.NewReader(f))
	}
}
