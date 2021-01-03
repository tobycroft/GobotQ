package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type SetNickNameRet struct {
	Ret string `json:"ret"`
}

func Setnickname(fromqq, nickname interface{}) (bool, error) {
	post := map[string]interface{}{
		"fromqq":   fromqq,
		"nickname": nickname,
	}
	data, err := Net.Post(app_conf.Http_Api+"/setnickname", nil, post, nil, nil)
	if err != nil {
		return false, err
	}
	var ret SetNickNameRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret)
	if err != nil {
		return false, err
	}
	if ret.Ret == "true" {
		return true, nil
	} else {
		return false, nil
	}
}
