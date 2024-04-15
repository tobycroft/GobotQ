package Aigc

import (
	"errors"
	"fmt"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/SystemParamModel"
	"strings"
	"time"
)

func Aigc_bing_text(text string) (AigcStruct, error) {
	post := Net.Net{}.New().SetTimeOut(100 * time.Second).SetUrl("http://aigc.aerofsx.com:84/v1/aigc/bing/text").
		SetPostData(map[string]string{"text": text}).
		SetHeader(map[string]string{"token": Calc.Any2String(SystemParamModel.Api_value("aigc"))}).
		PostFormData()
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
