package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/tuuz/Net"
)

type GetNickNameRet struct {
	Ret string `json:"ret"`
}

func Getnickname(fromqq, toqq, fromcache interface{}) (GetNickNameRet, error) {
	post := map[string]interface{}{
		"fromqq":    fromqq,
		"toqq":      toqq,
		"fromcache": fromcache,
	}
	data, err := Net.Post(botinfo["url"].(string)+"/getnickname", nil, post, nil, nil)
	if err != nil {
		return GetNickNameRet{}, err
	}
	var ret GetNickNameRet
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &ret)
	if err != nil {
		return GetNickNameRet{}, err
	}
	return ret, nil
}
