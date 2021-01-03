package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type ExitGroupRet struct {
	Ret string `json:"ret"`
}

func Exitgroup(fromqq, group interface{}) (bool, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"group":  group,
	}
	data, err := Net.Post(app_conf.Http_Api+"/exitgroup", nil, post, nil, nil)
	if err != nil {
		return false, err
	}
	var ret ExitGroupRet
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
