package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

/*
{
    "ret": "{\"retcode\":0,\"retmsg\":\"\",\"time\":\"1609680793\"}",
    "random": 1609702875,
    "req": 22503
}
*/

type SendGroupTempMsg struct {
	Ret    string `json:"ret"`
	Random int    `json:"random"`
	Req    int    `json:"req"`
}

type SendGroupTempMsgRet struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Time    string `json:"time"`
}

func Sendgrouptempmsg(fromqq, togroup, toqq, text interface{}) (SendGroupTempMsg, SendGroupTempMsgRet, error) {
	post := map[string]interface{}{
		"fromqq":  fromqq,
		"togroup": togroup,
		"toqq":    toqq,
		"text":    text,
	}
	data, err := Net.Post(app_conf.Http_Api+"/sendSendGroupTempMsg", nil, post, nil, nil)
	if err != nil {
		return SendGroupTempMsg{}, SendGroupTempMsgRet{}, err
	}
	var ret1 SendGroupTempMsg
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret1)
	if err != nil {
		return SendGroupTempMsg{}, SendGroupTempMsgRet{}, err
	}
	var ret2 SendGroupTempMsgRet
	err = jsr.UnmarshalFromString(ret1.Ret, &ret2)
	if err != nil {
		return ret1, SendGroupTempMsgRet{}, err
	}
	if ret2.Retcode != 0 {
		return ret1, ret2, errors.New(ret2.Retmsg)
	}
	return ret1, ret2, nil
}
