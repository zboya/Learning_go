package main

import (
	"time"
	"fmt"
	"encoding/base64"

	"net/http"
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
	"./youtu"
)



type detectFaceReq struct {
	AppID string    `json:"app_id"`         
	Image string    `json:"image"`   //base64编码的二进制图片数据       
	Mode  int 		`json:"mode"` //检测模式 0/1 正常/大脸模式
}

//Face 脸参数
type Face struct {
	FaceID     string  `json:"face_id"`    //人脸标识
	X          int32   `json:"x"`          //人脸框左上角x
	Y          int32   `json:"y"`          //人脸框左上角y
	Width      float32 `json:"width"`      //人脸框宽度
	Height     float32 `json:"height"`     //人脸框高度
	Gender     int32   `json:"gender"`     //性别 [0/(female)~100(male)]
	Age        int32   `json:"age"`        //年龄 [0~100]
	Expression int32   `json:"expression"` //object 	微笑[0(normal)~50(smile)~100(laugh)]
	Glass      bool    `json:"glass"`      //是否有眼镜 [true,false]
	Pitch      int32   `json:"pitch"`      //上下偏移[-30,30]
	Yaw        int32   `json:"yaw"`        //左右偏移[-30,30]
	Roll       int32   `json:"roll"`       //平面旋转[-180,180]
}
//DetectFaceRsp 脸检测返回
type DetectFaceRsp struct {
	SessionID   string `json:"session_id"`   //相应请求的session标识符，可用于结果查询
	ImageID     string `json:"image_id"`     //系统中的图片标识符，用于标识用户请求中的图片
	ImageWidth  int32  `json:"image_width"`  //请求图片的宽度
	ImageHeight int32  `json:"image_height"` //请求图片的高度
	Face        []Face `json:"face"`         //被检测出的人脸Face的列表
	ErrorCode   int    `json:"errorcode"`    //返回状态值
	ErrorMsg    string `json:"errormsg"`     //返回错误消息
}


const (
	normalMode int=iota
	bigFaceMode
)

func main() {
	p,b64auth:=youtu.sign()
	println("plain_text:",p,"base64:",b64auth)

	//Post value to api
	AppID:="1006405"
	pic_path:="./imageA.jpg"
	bImage,err:=ioutil.ReadFile(pic_path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "ReadFile() failed: %s\n", err)
		return
	}

	b64Image:=base64.StdEncoding.EncodeToString(bImage)
	Mode:=normalMode

	detect_req:=&detectFaceReq{AppID,b64Image,Mode}
	// println(detect_req)
	req,_:=json.Marshal(detect_req)
	body_req:=string(req)
	// println("json body_req:",body_req)
	detect_url:="http://api.youtu.qq.com/youtu/api/detectface"

	//http post detectfafe api
	client := &http.Client{
		Timeout: time.Duration(5 * time.Second),
	}
	httpreq, err := http.NewRequest("POST", detect_url, strings.NewReader(body_req))
	if err != nil {
		return
	}
	
	httpreq.Header.Add("Authorization", b64auth)
	httpreq.Header.Add("Content-Type", "text/json")
	httpreq.Header.Add("User-Agent", "")
	httpreq.Header.Add("Accept", "*/*")
	httpreq.Header.Add("Expect", "100-continue")
	resp, err := client.Do(httpreq)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	detect_rsp, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("read http respond failed")
		return
	}
	var json_detect_rsp DetectFaceRsp
	err = json.Unmarshal(detect_rsp,&json_detect_rsp )
	fmt.Printf("http rsp: %#v\n", json_detect_rsp)


}

