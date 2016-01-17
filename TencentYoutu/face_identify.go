package main

import (
	"fmt"
	"io/ioutil"
	"os"
	// "strconv"
	"./Tyoutu"
)

type PicData struct {
	name string
	pics []string
}


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
	picPath:="/home/bao/bao/work/listome/facedata/pic"

	pic_datas:=getPicData(picPath)
	for _,pic_data:=range pic_datas {
		
		personName:=pic_data.name
		pic_names:=pic_data.pics
		var pic_last string

		for _,pic_sig:=range pic_names {
			pic_last=pic_sig
		}
		println(pic_last)
		path:=picPath+"/"+personName+"/"+pic_last
		println(path)
		bimage,err:=ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ReadFile() failed: %s\n", err)
			return
		}
		rsp, err := yt.FaceIdentify(groupID,bimage)
		if err != nil {
			fmt.Printf("FaceIdentify failed: %s\n", err)
			return
		}
		fmt.Printf("rsp %#v\n", rsp)
	}
	

}

func getPicData(pic_path string) (pic_datas []PicData) {

	pic_datas = make([]PicData, 0, 10)
	pic_data:=new(PicData)
	
	dir,err:=ioutil.ReadDir(pic_path)

	if err != nil {
		println("ReadDir faild")
		return
	} 
	i:=0
	for _,v:= range dir {
		if i>=100 {break}
		if v.IsDir() {
			i++
			pic_data.pics=make([]string,0,10)
			pic_data.name =v.Name()
			
			subdir,_:=ioutil.ReadDir(pic_path+"/"+v.Name())
			for _,f := range subdir {
				// println(f.Name())
				pic_data.pics=append(pic_data.pics,f.Name())
			}

			pic_datas=append(pic_datas,*pic_data)
		}
	}
	return

}