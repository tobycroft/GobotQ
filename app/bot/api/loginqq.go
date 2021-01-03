package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type LoginqqRet struct {
	Ret string `json:"ret"`
}

func Loginqq(logonqq interface{}) (bool, error) {
	post := map[string]interface{}{
		"logonqq": logonqq,
	}
	data, err := Net.Post(app_conf.Http_Api+"/loginqq", nil, post, nil, nil)
	if err != nil {
		return false, err
	}
	var ret LoginqqRet
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
