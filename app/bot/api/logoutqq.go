package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type LogoutqqRet struct {
	Ret string `json:"ret"`
}

func Logoutqq(logonqq interface{}) (bool, error) {
	post := map[string]interface{}{
		"logonqq": logonqq,
	}
	data, err := Net.Post(app_conf.Http_Api+"/logoutqq", nil, post, nil, nil)
	if err != nil {
		return false, err
	}
	var ret LogoutqqRet
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
