package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type SetFriendAddRequest struct {
	Ret string `json:"ret"`
}

const Request_friend_denide = 2
const Request_friend_approve = 1

func Setfriendaddrequest(fromqq, group, qq, seq, Request_group_, Type, reason interface{}) (bool, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"group":  group,
		"qq":     qq,
		"seq":    seq,
		"op":     Request_group_,
		"type":   Type,
		"reason": reason,
	}
	data, err := Net.Post(app_conf.Http_Api+"/setnickname", nil, post, nil, nil)
	if err != nil {
		return false, err
	}
	var ret SetFriendAddRequest
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
