package main

import (
	"time"
	"math/rand"
	"fmt"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"

	"encoding/json"
)

//conf

type Auth struct {	
	AppID int
	SecretID string
	SecretKey string
	UserID string
	// Expired int
	// Rdm int
	// Now int64
}
const expired = 1000

/*AppID:="1006405"
SecretID:="AKIDPdd5DfYl0oeuelGdvoBCECgIZ6En2u1i"
SecretKey:="BxzXA7JTQTuB2QsxeqsglKLx6QoGK6kG"*/

func main() {
	p,b:=sign()
	println("plain_text:",p,"base64:",b)

}
/*
func str(n int64) (s string) {
	s=strconv.FormatInt(n,10)
	return
}*/

func sign() (plain_text,b64 string) {
	now:=time.Now().Unix()
	rdm:=rand.Intn(999999999)

	auth:=&Auth{1006405,"AKIDPdd5DfYl0oeuelGdvoBCECgIZ6En2u1i","BxzXA7JTQTuB2QsxeqsglKLx6QoGK6kG","1421967301"}
	plain_text = fmt.Sprintf("a=%d&k=%s&e=%d&t=%d&r=%d&u=%s&f=",
		auth.AppID,
		auth.SecretID,
		now+expired,
		now,
		rdm,
		auth.UserID)

	// println("plain_text:",plain_text)	
	h := hmac.New(sha1.New, []byte(auth.SecretKey))
	h.Write([]byte(plain_text))
	hm := h.Sum(nil)
	//attach orig_sign to hm
	dstSign := []byte(string(hm) + plain_text)
	b64 = base64.StdEncoding.EncodeToString(dstSign)
	return
	// println("base64:",b64)
}
