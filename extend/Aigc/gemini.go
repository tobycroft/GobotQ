package Aigc

import (
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/SystemParamModel"
	"time"
)

func Aigc_gemini_text(text string) (AigcStruct, error) {
	post := Net.Post{}.SetTimeOut(100*time.Second).PostUrlXEncode("http://10.0.0.182:84/v1/aigc/gemini/text", map[string]interface{}{
		"token": SystemParamModel.Api_value("aigc"),
	}, map[string]interface{}{
		"text": text,
	}, map[string]string{}, map[string]string{})
	var ag AigcStruct
	err := post.RetJson(&ag)
	if err != nil {
		return AigcStruct{}, err
	}
	return ag, err
}

type AigcStruct struct {
	Code int    `json:"code"`
	Echo string `json:"echo"`
}
