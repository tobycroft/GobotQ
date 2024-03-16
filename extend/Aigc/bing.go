package Aigc

import (
	"errors"
	"fmt"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/SystemParamModel"
	"strings"
	"time"
)

func Aigc_bing_text(text string) (AigcStruct, error) {
	post := Net.Post{}.New().SetTimeOut(100*time.Second).PostFormData("http://10.0.0.182:84/v1/aigc/bing/text", map[string]interface{}{
		"token": SystemParamModel.Api_value("aigc"),
	}, map[string]string{
		"text": text,
	}, map[string]string{}, map[string]string{})
	var ag AigcStruct
	err := post.RetJson(&ag)
	if err != nil {
		return AigcStruct{}, err
	}
	if ag.Code != 0 {
		return ag, errors.New("回答生成失败，请重新提问")
	}
	ag.Echo = strings.ReplaceAll(ag.Echo, "<br>", "\r")
	fmt.Println("AIGC_RET", ag.Echo)
	return ag, err
}
