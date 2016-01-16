package youtu

import (
	"time"
	"math/rand"
	"fmt"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
)

//conf

type Auth struct {	
	AppID string
	SecretID string
	SecretKey string
	UserID string
	// Expired int
	// Rdm int
	// Now int64
}

const expired = 1000
var debug bool = false


func Sign() (plain_text,b64 string) {
	//conf of sign
	AppID:="1006405"
	SecretID:="AKIDPdd5DfYl0oeuelGdvoBCECgIZ6En2u1i"
	SecretKey:="BxzXA7JTQTuB2QsxeqsglKLx6QoGK6kG"
	UserID:="1421967301"

	now:=time.Now().Unix()
	rdm:=rand.Intn(999999999)

	auth:=&Auth{AppID,SecretID,SecretKey,UserID}
	plain_text = fmt.Sprintf("a=%s&k=%s&e=%d&t=%d&r=%d&u=%s&f=",
		auth.AppID,
		auth.SecretID,
		now+expired,
		now,
		rdm,
		auth.UserID)

	if debug {println("plain_text:",plain_text)}

	h := hmac.New(sha1.New, []byte(auth.SecretKey))
	h.Write([]byte(plain_text))
	hm := h.Sum(nil)
	//attach orig_sign to hm
	dstSign := []byte(string(hm) + plain_text)
	b64 = base64.StdEncoding.EncodeToString(dstSign)
	if debug {println("base64:",b64)}
	return
}
