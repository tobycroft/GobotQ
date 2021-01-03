package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type SetNickNameRet struct {
	Ret string `json:"ret"`
}

func Setnickname(fromqq, toqq, random, req, time interface{}) (SetNickNameRet, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"toqq":   toqq,
		"random": random,
		"req":    req,
		"time":   time,
	}
	data, err := Net.Post(app_conf.Http_Api+"/setnickname", nil, post, nil, nil)
	if err != nil {
		return SetNickNameRet{}, err
	}
	var dpmr SetNickNameRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &dpmr)
	if err != nil {
		return SetNickNameRet{}, err
	}
	return dpmr, nil
}
