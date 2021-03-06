package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type Retract_private struct {
	Fromqq interface{}
	Toqq   interface{}
	Random interface{}
	Req    interface{}
	Time   interface{}
}

var Retract_chan_private = make(chan Retract_private, 20)
var Retract_chan_private_instant = make(chan Retract_private, 20)

type DeletePrivateMsgRet struct {
	Ret string `json:"ret"`
}

func Deleteprivatemsg(fromqq, toqq, random, req, time interface{}) (bool, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"toqq":   toqq,
		"random": random,
		"req":    req,
		"time":   time,
	}
	data, err := Net.Post(app_conf.Http_Api+"/deleteprivatemsg", nil, post, nil, nil)
	if err != nil {
		return false, err
	}
	var ret DeletePrivateMsgRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret)
	if err != nil {
		return false, err
	}
	if ret.Ret == "true" {
		return true, nil
	} else {
		return false, nil
	}
}
