package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type UploadFacePic struct {
	Ret string `json:"ret"`
}

type UploadFacePicRet struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Time    string `json:"time"`
}

func Uploadfacepic(fromqq, b64 interface{}) (UploadFacePic, UploadFacePicRet, error) {
	post := map[string]interface{}{
		"fromqq":   fromqq,
		"fromtype": 0,
		"pic":      b64,
	}
	data, err := Net.Post(app_conf.Http_Api+"/UploadFacePic", nil, post, nil, nil)
	if err != nil {
		return UploadFacePic{}, UploadFacePicRet{}, err
	}
	var ret1 UploadFacePic
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret1)
	if err != nil {
		return UploadFacePic{}, UploadFacePicRet{}, err
	}
	var ret2 UploadFacePicRet
	err = jsr.UnmarshalFromString(ret1.Ret, &ret2)
	if err != nil {
		return ret1, UploadFacePicRet{}, err
	}
	if ret2.Retcode != 0 {
		return ret1, ret2, errors.New(ret2.Retmsg)
	}
	return ret1, ret2, nil
}
