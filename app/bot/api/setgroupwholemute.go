package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type SetGroupWholeMuteRet struct {
	Ret string `json:"ret"`
}

func Setgroupwholemute(fromqq, group interface{}, ismute bool) (bool, error) {
	post := map[string]interface{}{
		"fromqq": fromqq,
		"group":  group,
		"ismute": ismute,
	}
	data, err := Net.Post(botinfo["url"].(string)+"/setgroupwholemute", nil, post, nil, nil)
	if err != nil {
		return false, err
	}
	var ret SetGroupWholeMuteRet
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
