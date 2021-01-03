package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type AddGroup struct {
	Ret string `json:"ret"`
}

type AddGroupRet struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Time    string `json:"time"`
}

func Addgroup(fromqq, togroup, text interface{}) (AddGroup, AddGroupRet, error) {
	post := map[string]interface{}{
		"fromqq":  fromqq,
		"togroup": togroup,
		"text":    text,
	}
	data, err := Net.Post(app_conf.Http_Api+"/addgroup", nil, post, nil, nil)
	if err != nil {
		return AddGroup{}, AddGroupRet{}, err
	}
	var ret1 AddGroup
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret1)
	if err != nil {
		return AddGroup{}, AddGroupRet{}, err
	}
	var ret2 AddGroupRet
	err = jsr.UnmarshalFromString(ret1.Ret, &ret2)
	if err != nil {
		return ret1, AddGroupRet{}, err
	}
	if ret2.Retcode != 0 {
		return ret1, ret2, errors.New(ret2.Retmsg)
	}
	return ret1, ret2, nil
}
