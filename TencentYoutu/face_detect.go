package main

import (
	"time"
	"math/rand"
	"strconv"
	"fmt"
)

//conf

func main() {
	AppID:="1006405"
	SecretID:="AKIDPdd5DfYl0oeuelGdvoBCECgIZ6En2u1i"
	SecretKey:="BxzXA7JTQTuB2QsxeqsglKLx6QoGK6kG"
	
	now:=time.Now().Unix()
	println(now)
	rdm:=rand.Intn(999999999)
	expired := 0
	qq:="1421967301"
	fmt.Printf("%T rdm: %d ",rdm,rdm)
	plain_text:="a=" + AppID + "&k=" + SecretID + "&e=" + str((int64)expired) + "&t=" + str(now) + "&r=" + str((int64)rdm) + "&u=" + qq + "&f="
	println("plain_text:",plain_text)	
}

func str(n int64) (s string) {
	s=strconv.FormatInt(n)
	return
}
