package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

/*
{
    "ret": "{\"retcode\":0,\"retmsg\":\"\",\"time\":\"1609564779\"}"
}
*/

type GroupMsg struct {
	Ret string `json:"ret"`
}

type GroupMsgRet struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Time    string `json:"time"`
}

func Sendgroupmsg(fromqq, togroup interface{}, text string) (GroupMsg, GroupMsgRet, error) {
	post := map[string]interface{}{
		"fromqq":  fromqq,
		"togroup": togroup,
		"text":    text,
	}
	data, err := Net.Post(app_conf.Http_Api+"/sendgroupmsg", nil, post, nil, nil)
	if err != nil {
		return GroupMsg{}, GroupMsgRet{}, err
	}
	var gm GroupMsg
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gm)
	if err != nil {
		return GroupMsg{}, GroupMsgRet{}, err
	}
	var ret GroupMsgRet
	err = jsr.UnmarshalFromString(gm.Ret, &ret)
	if err != nil {
		return gm, GroupMsgRet{}, err
	}
	if ret.Retcode != 0 {
		return gm, ret, errors.New(ret.Retmsg)
	}
	return gm, ret, nil
}
