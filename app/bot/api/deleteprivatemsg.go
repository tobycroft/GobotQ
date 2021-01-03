package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type DeletePrivateMsgRet struct {
	Ret string `json:"ret"`
}

func Deleteprivatemsg(fromqq, toqq, random, req, time interface{}) (DeletePrivateMsgRet, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"toqq":   toqq,
		"random": random,
		"req":    req,
		"time":   time,
	}
	data, err := Net.Post(app_conf.Http_Api+"/deleteprivatemsg", nil, post, nil, nil)
	if err != nil {
		return DeletePrivateMsgRet{}, err
	}
	var dpmr DeletePrivateMsgRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &dpmr)
	if err != nil {
		return DeletePrivateMsgRet{}, err
	}
	return dpmr, nil
}
