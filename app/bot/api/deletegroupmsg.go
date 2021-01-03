package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type DeleteGroupMsgRet struct {
	Ret string `json:"ret"`
}

func Deletegroupmsg(fromqq, group, random, req interface{}) (DeleteGroupMsgRet, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"group":  group,
		"random": random,
		"req":    req,
	}
	data, err := Net.Post(app_conf.Http_Api+"/deletegroupmsg", nil, post, nil, nil)
	if err != nil {
		return DeleteGroupMsgRet{}, err
	}
	var dgmr DeleteGroupMsgRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &dgmr)
	if err != nil {
		return DeleteGroupMsgRet{}, err
	}
	return dgmr, nil
}
