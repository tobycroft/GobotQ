package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

/*
{
    "ret": "{\"retcode\":0,\"retmsg\":\"\",\"time\":\"1609562723\"}",
    "random": 1609583468,
    "req": 21315
}
*/

type PrivateMsg struct {
	Ret    string `json:"ret"`
	Random int    `json:"random"`
	Req    int    `json:"req"`
}

type PrivateMsgRet struct {
	Retcode int    `json:"retcode"`
	Retmsg  string `json:"retmsg"`
	Time    string `json:"time"`
}

func Sendprivatemsg(fromqq, toqq interface{}, text string) (PrivateMsg, PrivateMsgRet, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"toqq":   toqq,
		"text":   text,
	}
	data, err := Net.Post(app_conf.Http_Api+"/sendprivatemsg", nil, post, nil, nil)
	if err != nil {
		return PrivateMsg{}, PrivateMsgRet{}, err
	}
	var pm PrivateMsg
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &pm)
	if err != nil {
		return PrivateMsg{}, PrivateMsgRet{}, err
	}
	var ret PrivateMsgRet
	err = jsr.UnmarshalFromString(pm.Ret, &ret)
	if err != nil {
		return pm, PrivateMsgRet{}, err
	}
	if ret.Retcode != 0 {
		return pm, ret, errors.New(ret.Retmsg)
	}
	return pm, ret, nil
}
