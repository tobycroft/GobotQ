package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type DeleteGroupMsgRet struct {
	Ret string `json:"ret"`
}

type Retract_group struct {
	Fromqq interface{}
	Group  interface{}
	Random interface{}
	Req    interface{}
}

var Retract_chan_group = make(chan Retract_group, 20)
var Retract_chan_group_instant = make(chan Retract_group, 20)

func Deletegroupmsg(fromqq, group, random, req interface{}) (bool, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"group":  group,
		"random": random,
		"req":    req,
	}
	data, err := Net.Post(app_conf.Http_Api+"/deletegroupmsg", nil, post, nil, nil)
	if err != nil {
		return false, err
	}
	var ret DeleteGroupMsgRet
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
