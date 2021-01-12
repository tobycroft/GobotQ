package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type SendPrivatePicRet struct {
	Ret string `json:"ret"`
}

func sendprivatepic_file(fromqq, toqq, path interface{}) (string, error) {
	post := map[string]interface{}{
		"fromqq":   fromqq,
		"toqq":     toqq,
		"fromtype": 1,
		"path":     path,
	}
	data, err := Net.Post(app_conf.Http_Api+"/sendprivatepic", nil, post, nil, nil)
	if err != nil {
		return "", err
	}
	var gm GroupMsg
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gm)
	if err != nil {
		return "", err
	}
	var ret SendPrivatePicRet
	err = jsr.UnmarshalFromString(gm.Ret, &ret)
	if err != nil {
		return "", err
	}
	return ret.Ret, nil
}

func sendprivatepic_base64(fromqq, toqq, pic interface{}) (string, error) {
	post := map[string]interface{}{
		"fromqq":   fromqq,
		"toqq":     toqq,
		"fromtype": 0,
		"pic":      pic,
	}
	data, err := Net.Post(app_conf.Http_Api+"/sendprivatepic", nil, post, nil, nil)
	if err != nil {
		return "", err
	}
	var gm GroupMsg
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gm)
	if err != nil {
		return "", err
	}
	var ret SendPrivatePicRet
	err = jsr.UnmarshalFromString(gm.Ret, &ret)
	if err != nil {
		return "", err
	}
	return ret.Ret, nil
}

func sendprivatepic_remote(fromqq, toqq, url interface{}) (string, error) {
	post := map[string]interface{}{
		"fromqq":   fromqq,
		"toqq":     toqq,
		"fromtype": 2,
		"url":      url,
	}
	data, err := Net.Post(app_conf.Http_Api+"/sendprivatepic", nil, post, nil, nil)
	if err != nil {
		return "", err
	}
	var gm GroupMsg
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gm)
	if err != nil {
		return "", err
	}
	var ret SendPrivatePicRet
	err = jsr.UnmarshalFromString(gm.Ret, &ret)
	if err != nil {
		return "", err
	}
	return ret.Ret, nil
}
