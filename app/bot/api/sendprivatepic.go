package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
	"strings"
)

type SendPrivatePicRet struct {
	Ret string `json:"ret"`
}

type SendPrivatePic struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Time    string `json:"time"`
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
	if strings.Contains(ret.Ret, "pic,hash") {
		return ret.Ret, nil
	} else {
		var ret2 SendPrivatePic
		err = jsr.UnmarshalFromString(ret.Ret, &ret2)
		if err != nil {
			return "", err
		} else {
			return "", errors.New(ret2.Retmsg)
		}
	}
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
	if strings.Contains(ret.Ret, "pic,hash") {
		return ret.Ret, nil
	} else {
		var ret2 SendPrivatePic
		err = jsr.UnmarshalFromString(ret.Ret, &ret2)
		if err != nil {
			return "", err
		} else {
			return "", errors.New(ret2.Retmsg)
		}
	}
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
	if strings.Contains(ret.Ret, "pic,hash") {
		return ret.Ret, nil
	} else {
		var ret2 SendPrivatePic
		err = jsr.UnmarshalFromString(ret.Ret, &ret2)
		if err != nil {
			return "", err
		} else {
			return "", errors.New(ret2.Retmsg)
		}
	}
}
