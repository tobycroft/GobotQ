package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type KickGroupMemberRet struct {
	Ret string `json:"ret"`
}

func Kickgroupmember(fromqq, toqq, random, req, time interface{}) (KickGroupMemberRet, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"toqq":   toqq,
		"random": random,
		"req":    req,
		"time":   time,
	}
	data, err := Net.Post(app_conf.Http_Api+"/kickgroupmember", nil, post, nil, nil)
	if err != nil {
		return KickGroupMemberRet{}, err
	}
	var ret KickGroupMemberRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret)
	if err != nil {
		return KickGroupMemberRet{}, err
	}
	return ret, nil
}
