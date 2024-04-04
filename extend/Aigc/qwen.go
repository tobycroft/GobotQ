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

func Aigc_qwen_base(message string) (AigcStruct, error) {
	post := Net.Net{}.New().SetTimeOut(100 * time.Second).SetUrl("http://10.0.0.182:84/v1/qwen/api/raw").
		SetPostData(map[string]string{"message": message, "chat_id": "111"}).
		SetHeader(map[string]string{"Authorization": "Bearer " + Calc.Any2String(SystemParamModel.Api_value("subtoken"))}).
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
