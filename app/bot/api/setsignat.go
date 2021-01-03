package api

import (
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type SetSigNat struct {
	Ret string `json:"ret"`
}

func Setsignat(fromqq, signature interface{}) (SetSigNat, error) {
	post := map[string]interface{}{
		"fromqq":    fromqq,
		"signature": signature,
	}
	data, err := Net.Post(app_conf.Http_Api+"/setsignat", nil, post, nil, nil)
	if err != nil {
		return SetSigNat{}, err
	}
	var dpmr SetSigNat
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &dpmr)
	if err != nil {
		return SetSigNat{}, err
	}
	return dpmr, nil
}
