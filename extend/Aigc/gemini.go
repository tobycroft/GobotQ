package Aigc

import (
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/SystemParamModel"
)

func Aigc_gemini_text(text string) (AigcGeminiStruct, error) {
	post := Net.Post{}.PostUrlXEncode("http://aigc.aerofsx.com:81/v1/aigc/gemini/text", map[string]interface{}{
		"token": SystemParamModel.Api_value("aigc"),
	}, map[string]interface{}{
		"text": text,
	}, map[string]string{}, map[string]string{})
	var ag AigcGeminiStruct
	err := post.RetJson(&ag)
	if err != nil {
		return AigcGeminiStruct{}, err
	}
	return ag, err
}

type AigcGeminiStruct struct {
	Code int `json:"code"`
	Data struct {
		Image []string `json:"image"`
		Text  string   `json:"text"`
	} `json:"data"`
	Echo string `json:"echo"`
}
