package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type KickGroupMemberRet struct {
	Ret string `json:"ret"`
}

func Kickgroupmember(fromqq, group, toqq interface{}, ignoreaddgrequest bool) (bool, error) {
	post := map[string]interface{}{
		"fromqq":            fromqq,
		"group":             group,
		"toqq":              toqq,
		"ignoreaddgrequest": ignoreaddgrequest,
	}
	data, err := Net.Post(app_conf.Http_Api+"/kickgroupmember", nil, post, nil, nil)
	if err != nil {
		return KickGroupMemberRet{}, err
	}
	var ret KickGroupMemberRet
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
