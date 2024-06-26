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

func Aigc_qwen(chat_id, message string) (AigcEchoStruct, error) {
	post := Net.Net{}.New().SetTimeOut(100 * time.Second).SetUrl("http://aigc.aerofsx.com:81/v1/qwen/api/raw").
		SetPostData(map[string]string{"message": message, "chat_id": chat_id}).
		SetHeader(map[string]string{"Authorization": "Bearer " + Calc.Any2String(SystemParamModel.Api_value("subtoken"))}).
		PostFormData()
	var ag AigcEchoStruct
	err := post.RetJson(&ag)
	if err != nil {
		return AigcEchoStruct{}, err
	}
	if ag.Code != 0 {
		return ag, errors.New("回答生成失败，请重新提问")
	}
	ag.Echo = strings.ReplaceAll(ag.Echo, "<br>", "\r")
	fmt.Println("AIGC_RET", ag.Echo)
	return ag, err
}
