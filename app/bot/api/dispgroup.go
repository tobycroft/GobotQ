package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type DispGroupRet struct {
	Ret string `json:"ret"`
}

func Dispgroup(fromqq, group interface{}) (bool, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"group":  group,
	}
	data, err := Net.Post(botinfo["url"].(string)+"/dispgroup", nil, post, nil, nil)
	if err != nil {
		return false, err
	}
	var ret DispGroupRet
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
