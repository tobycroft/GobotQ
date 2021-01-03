package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type SetSigNat struct {
	Ret string `json:"ret"`
}

func Setsignat(fromqq, signature interface{}) (bool, error) {
	post := map[string]interface{}{
		"fromqq":    fromqq,
		"signature": signature,
	}
	data, err := Net.Post(app_conf.Http_Api+"/setsignat", nil, post, nil, nil)
	if err != nil {
		return false, err
	}
	var ret SetSigNat
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
