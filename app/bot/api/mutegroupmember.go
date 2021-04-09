package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type MuteGroupmeMberRet struct {
	Ret string `json:"ret"`
}

func Mutegroupmember(fromqq, group, toqq interface{}, time float64) (bool, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"group":  group,
		"toqq":   toqq,
		"time":   time,
	}
	data, err := Net.Post(botinfo["url"].(string)+"/mutegroupmember", nil, post, nil, nil)
	if err != nil {
		return false, err
	}
	var ret MuteGroupmeMberRet
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
