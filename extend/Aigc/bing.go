package Aigc

import (
	"fmt"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/SystemParamModel"
	"strings"
	"time"
)

func Aigc_bing_text(text string) (AigcStruct, error) {
	post := Net.Post{}.New().SetTimeOut(100*time.Second).PostUrlXEncode("http://10.0.0.182:84/v1/aigc/bing/text", map[string]interface{}{
		"token": SystemParamModel.Api_value("aigc"),
	}, map[string]interface{}{
		"text": text,
	}, map[string]string{}, map[string]string{})
	var ag AigcStruct
	err := post.RetJson(&ag)
	if err != nil {
		return AigcStruct{}, err
	}
	ag.Echo = strings.ReplaceAll(ag.Echo, "<br>", "\r")
	fmt.Println("AIGC_RET", ag.Echo)
	return ag, err
}
