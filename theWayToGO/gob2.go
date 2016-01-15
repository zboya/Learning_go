package main

import (
    "encoding/gob"
//    "log"
    "os"
	"fmt"
)

type Address struct {
    Type             string
    City             string
    Country          string
}

type VCard struct {
    FirstName   string
    LastName    string
    Addresses   []*Address
    Remark      string
}

var content string

func main() {
	pa := &Address{"private", "Aartselaar","Belgium"}
    wa := &Address{"work", "Boom", "Belgium"}
    vc := VCard{"Jan", "Kersschot", []*Address{pa,wa}, "none"}	
	
	file,_:=os.OpenFile("vcard.gob",os.O_CREATE|os.O_WRONLY,0666)
	defer file.Close()
	enc:=gob.NewEncoder(file)
	enc.Encode(vc)

	//read and decode
	file,_=os.Open("vcard.gob")
	var vcard VCard
	defer file.Close()
	dec:=gob.NewDecoder(file)
	dec.Decode(&vcard)

	fmt.Printf("%v\n",vcard)
	fmt.Printf("%s\n",vcard.FirstName)
	
}
