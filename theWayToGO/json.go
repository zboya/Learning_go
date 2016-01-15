package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type string `json:"type"`
	City string `json:"city"`
	Country string `json:"country"`
}

type VCard struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Addresses []*Address `json:"addresses"`
	Remark string `json:"remark"`
}

func main () {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	js,_:=json.Marshal(vc)
	fmt.Printf("json format:%s\n",js)
	fmt.Printf("json type:%T\n",js)

	file,_:=os.OpenFile("vcad.json",os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()
	enc:=json.NewEncoder(file)
	err:=enc.Encode(vc)
	if err!=nil {
		log.Println("json encoding error")
	}
	
	var f interface{}
	err = json.Unmarshal(js,&f)
	fmt.Printf("decoding json:%v\n type:%T\n",f,f)
	if err!=nil {
		log.Println("json decoding error")
	}
	r := f.(map[string]interface{})
	println(r["first_name"].(string))
	fmt.Printf("==========%T\n",r["addresses"])	
	v := r["addresses"].([]interface{})

	for _,vv := range v {
		rr := vv.(map[string]interface{})
		println(rr["type"].(string))
	}

	vcard := new(VCard)

	err = json.Unmarshal(js,vcard)
	if err != nil {
		log.Println("Json unmarshal error:",err)
		return
	}
	
	for _,v := range vcard.Addresses{
		fmt.Println(v.Type)
	}

	fmt.Println(vcard.Addresses[0].Type)

}
