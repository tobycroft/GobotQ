package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
	"strings"
)

type SendGroupPicRet struct {
	Ret string `json:"ret"`
}

func sendgrouppic_file(fromqq, togroup, path interface{}) (string, error) {
	post := map[string]interface{}{
		"fromqq":   fromqq,
		"togroup":  togroup,
		"fromtype": 1,
		"path":     path,
	}
	data, err := Net.Post(app_conf.Http_Api+"/sendgrouppic", nil, post, nil, nil)
	if err != nil {
		return "", err
	}
	var gm GroupMsg
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gm)
	if err != nil {
		return "", err
	}
	var ret SendGroupPicRet
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

func sendgrouppic_base64(fromqq, togroup, pic interface{}) (string, error) {
	post := map[string]interface{}{
		"fromqq":   fromqq,
		"togroup":  togroup,
		"fromtype": 0,
		"pic":      pic,
	}
	data, err := Net.Post(app_conf.Http_Api+"/sendgrouppic", nil, post, nil, nil)
	if err != nil {
		return "", err
	}
	var gm GroupMsg
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gm)
	if err != nil {
		return "", err
	}
	var ret SendGroupPicRet
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

func sendgrouppic_remote(fromqq, togroup, url interface{}) (string, error) {
	post := map[string]interface{}{
		"fromqq":   fromqq,
		"togroup":  togroup,
		"fromtype": 2,
		"url":      url,
	}
	data, err := Net.Post(app_conf.Http_Api+"/sendgrouppic", nil, post, nil, nil)
	if err != nil {
		return "", err
	}
	var gm GroupMsg
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gm)
	if err != nil {
		return "", err
	}
	var ret SendGroupPicRet
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
