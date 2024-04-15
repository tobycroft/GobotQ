package Aigc

import (
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/SystemParamModel"
	"time"
)

func Aigc_gemini_text(text string) (AigcStruct, error) {
	post := Net.Net{}.New().SetTimeOut(100 * time.Second).SetUrl("http://aigc.aerofsx.com:84/v1/aigc/gemini/text").
		SetPostData(map[string]string{"text": text}).
		SetHeader(map[string]string{"token": Calc.Any2String(SystemParamModel.Api_value("aigc"))}).
		PostFormData()
	var ag AigcStruct
	err := post.RetJson(&ag)
	if err != nil {
		return AigcStruct{}, err
	}
	return ag, err
}

type AigcStruct struct {
	Code int    `json:"code"`
	Data string `json:"data"`
	Echo string `json:"echo"`
}

type AigcEchoStruct struct {
	Code int    `json:"code"`
	Echo string `json:"echo"`
}
