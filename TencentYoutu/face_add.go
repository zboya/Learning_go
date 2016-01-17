
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"./Tyoutu"
)

func main() {
	//Register your app on http://open.youtu.qq.com
	//Get the following details
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
	picPath:="/home/bao/bao/work/listome/facedata/pic"

	groupID:=make([]string,0,10)
	groupID=append(groupID,"yoututest")
	fmt.Printf("%v\n",groupID)
	AddPerson(picPath,yt,groupID)
	
}

type PicData struct {
	name string
	pics []string
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
//build newperson and newgroup
func AddPerson(pic_path string ,y *youtu.Youtu,groupIDs []string) {
	pic_datas:=getPicData(pic_path)
	i:=1
	for _,pic_data:=range pic_datas {
		personName:=pic_data.name
		pic_names:=pic_data.pics

		personID:="test"+strconv.Itoa(i)
		i++
		picPath:=pic_path+"/"+personName+"/"+pic_names[0]
		println(picPath)
		image,err:=ioutil.ReadFile(picPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "ReadFile() failed: %s\n", err)
			return
		}
		tag:="test"
		rsp, err :=y.NewPerson(personID, personName ,groupIDs, image , tag)
		if err != nil && rsp.ErrorMsg != "ERROR_PERSON_EXISTED" {
			println("NewPerson failed: ", err)
			return
		}
		fmt.Printf("rsp: %#v\n", rsp)
		n:=0
		for _,pics:=range pic_names {
			if n>10 {break}  //<=10 pics
			n++
			// fmt.Printf("pics: %#v\n", pics)
			picPath=pic_path+"/"+personName+"/"+pics
			println(picPath)
			bimage,err:=ioutil.ReadFile(picPath)
			bimages:=[][]byte{bimage}
			if err != nil {
				fmt.Fprintf(os.Stderr, "ReadFile() failed: %s\n", err)
				return
			}
			afrsp, err := y.AddFace(personID, bimages, tag)
			if err != nil {
				fmt.Fprintf(os.Stderr, "AddFace() failed: %s", err)
				return
			}
			fmt.Printf("rsp: %#v\n", afrsp)

		}

	}


}