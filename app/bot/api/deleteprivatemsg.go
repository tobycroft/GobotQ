package api

import (
	"errors"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

func Deleteprivatemsg(fromqq, toqq, random, req, time interface{}) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"toqq":   toqq,
		"random": random,
		"req":    req,
		"time":   time,
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
