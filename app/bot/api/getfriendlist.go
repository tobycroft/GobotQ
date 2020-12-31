package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type FriendList []struct {
	UIN      int    `json:"UIN"`
	NickName string `json:"NickName"`
	Remark   string `json:"Remark"`
	Email    string `json:"Email"`
}

func Getfriendlist(bot interface{}) (FriendList, error) {
	post := map[string]interface{}{
		"logonqq": bot,
	}
	data, err := Net.Post(app_conf.Http_Api+"/getfriendlist", nil, post, nil, nil)
	if err != nil {
		return nil, err
	}
	var fl FriendList
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &fl)
	if err != nil {
		return nil, err
	}
	return fl, nil
}
