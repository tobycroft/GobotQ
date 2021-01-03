package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type DeleteFriend struct {
	Ret string `json:"ret"`
}

type DeleteFriendRet struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Time    string `json:"time"`
}

func deletefriend() (DeleteFriend, DeleteFriendRet, error) {
	post := map[string]interface{}{
		"fromqq":  fromqq,
		"togroup": togroup,
		"text":    text,
	}
	data, err := Net.Post(app_conf.Http_Api+"/deleteFriend", nil, post, nil, nil)
	if err != nil {
		return DeleteFriend{}, DeleteFriendRet{}, err
	}
	var ret1 DeleteFriend
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret1)
	if err != nil {
		return DeleteFriend{}, DeleteFriendRet{}, err
	}
	var ret2 DeleteFriendRet
	err = jsr.UnmarshalFromString(ret1.Ret, &ret2)
	if err != nil {
		return ret1, DeleteFriendRet{}, err
	}
	if ret2.Retcode != 0 {
		return ret1, ret2, errors.New(ret2.Retmsg)
	}
	return ret1, ret2, nil
}
