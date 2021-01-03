package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type AddFriend struct {
	Ret string `json:"ret"`
}

type AddFriendRet struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Time    string `json:"time"`
}

func Addfriend(fromqq, toqq, text, remark interface{}) (AddFriend, AddFriendRet, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"toqq":   toqq,
		"text":   text,
		"remark": remark,
	}
	data, err := Net.Post(app_conf.Http_Api+"/setgroupcard", nil, post, nil, nil)
	if err != nil {
		return AddFriend{}, AddFriendRet{}, err
	}
	var ret1 AddFriend
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret1)
	if err != nil {
		return AddFriend{}, AddFriendRet{}, err
	}
	var ret2 AddFriendRet
	err = jsr.UnmarshalFromString(ret1.Ret, &ret2)
	if err != nil {
		return ret1, AddFriendRet{}, err
	}
	if ret2.Retcode != 0 {
		return ret1, ret2, errors.New(ret2.Retmsg)
	}
	return ret1, ret2, nil
}
