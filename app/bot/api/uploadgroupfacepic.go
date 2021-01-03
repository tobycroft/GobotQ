package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type UploadGroupFacePic struct {
	Ret string `json:"ret"`
}

type UploadGroupFacePicRet struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Time    string `json:"time"`
}

func Uploadgroupfacepic(fromqq, group, pic interface{}) (UploadGroupFacePic, UploadGroupFacePicRet, error) {
	post := map[string]interface{}{
		"fromqq":   fromqq,
		"group":    group,
		"fromtype": 0,
		"pic":      pic,
	}
	data, err := Net.Post(app_conf.Http_Api+"/UploadGroupFacePic", nil, post, nil, nil)
	if err != nil {
		return UploadGroupFacePic{}, UploadGroupFacePicRet{}, err
	}
	var ret1 UploadGroupFacePic
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret1)
	if err != nil {
		return UploadGroupFacePic{}, UploadGroupFacePicRet{}, err
	}
	var ret2 UploadGroupFacePicRet
	err = jsr.UnmarshalFromString(ret1.Ret, &ret2)
	if err != nil {
		return ret1, UploadGroupFacePicRet{}, err
	}
	if ret2.Retcode != 0 {
		return ret1, ret2, errors.New(ret2.Retmsg)
	}
	return ret1, ret2, nil
}
