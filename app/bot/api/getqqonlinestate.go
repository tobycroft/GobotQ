package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type GetqqOnlineStateRet struct {
	Ret string `json:"ret"`
}

func Getqqonlinestate(logonqq, qq interface{}) (string, error) {
	post := map[string]interface{}{
		"logonqq": logonqq,
		"qq":      qq,
	}
	data, err := Net.Post(botinfo["url"].(string)+"/getqqonlinestate", nil, post, nil, nil)
	if err != nil {
		return "", err
	}
	var ret GetqqOnlineStateRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret)
	if err != nil {
		return "", err
	}
	return ret.Ret, nil
}
