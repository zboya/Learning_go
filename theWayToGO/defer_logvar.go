package main

import (
	"io"
	"log"
) 

func f1(s string) (n int,err error) {
	defer func() {
		log.Printf("f1(%q)=%d,%v",s,n,err)
	}()
	return 7,io.EOF
}

func main(){
	f1("GO")
}
