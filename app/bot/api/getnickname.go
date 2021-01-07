package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type GetNickNameRet struct {
	Ret string `json:"ret"`
}

func Getnickname(fromqq, toqq, fromcache interface{}) (GetNickNameRet, error) {
	post := map[string]interface{}{
		"fromqq":    fromqq,
		"toqq":      toqq,
		"fromcache": fromcache,
	}
	data, err := Net.Post(app_conf.Http_Api+"/getnickname", nil, post, nil, nil)
	if err != nil {
		return GetNickNameRet{}, err
	}
	var ret GetNickNameRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret)
	if err != nil {
		return GetNickNameRet{}, err
	}
	return ret, nil
}
