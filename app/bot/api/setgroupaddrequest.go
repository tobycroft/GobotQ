package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type SetGroupAddRequestRet struct {
	Ret string `json:"ret"`
}

const Request_group_denide = 12
const Request_group_approve = 11

const Request_group_type_invite = 1
const Request_group_type_join = 3

func Setgroupaddrequest(fromqq, group, qq, seq, Request_group_, Request_group_type_, reason interface{}) (bool, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"group":  group,
		"qq":     qq,
		"seq":    seq,
		"op":     Request_group_,
		"type":   Request_group_type_,
		"reason": reason,
	}
	data, err := Net.Post(app_conf.Http_Api+"/setgroupaddrequest", nil, post, nil, nil)
	if err != nil {
		return false, err
	}
	var ret SetGroupAddRequestRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret)
	if err != nil {
		return false, err
	}
	if ret.Ret == "OK" {
		return true, nil
	} else {
		return false, nil
	}
}
