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

func Setfriendaddrequest(fromqq, qq, seq, Request_friend_ interface{}) (bool, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"qq":     qq,
		"seq":    seq,
		"op":     Request_friend_,
	}
	data, err := Net.Post(botinfo["url"].(string)+"/setfriendaddrequest", nil, post, nil, nil)
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
