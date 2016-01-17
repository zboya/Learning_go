package main

import (
	"fmt"
	// "io/ioutil"
	"os"
	"strconv"
	"./Tyoutu"
)

func main() {
	appID := uint32(1006405)
	secretID := "AKIDPdd5DfYl0oeuelGdvoBCECgIZ6En2u1i"
	secretKey := "BxzXA7JTQTuB2QsxeqsglKLx6QoGK6kG"
	userID := "1421967301"


	as, err := youtu.NewAppSign(appID, secretID, secretKey, userID)
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewAppSign() failed: %s\n", err)
		return
	}

	yt := youtu.Init(as, youtu.DefaultHost)

	//
	groupID:="yoututest"
	rsp, err := yt.GetPersonIDs(groupID)
		if err != nil {
			fmt.Printf("GetPersonIDs failed: %s\n", err)
			return
		}
		fmt.Printf("rsp %#v\n", rsp)

	for i := 1; i <=100; i++ {
		personID:="test"+strconv.Itoa(i)
	
		rsp, err := yt.GetInfo(personID)
		if err != nil {
			fmt.Printf("GetInfo failed: %s\n", err)
			return
		}
		fmt.Printf("rsp %#v\n", rsp)
	}
	

}

