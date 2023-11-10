package service

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/GroupAutoReplyModel"

	"strings"
)

func Serv_auto_reply(gid interface{}, text string) (string, bool) {
	rand := Calc.Rand(1, 99)
	data := GroupAutoReplyModel.Api_find(gid, text, rand)
	if len(data) > 0 {
		return Calc.Any2String(data["value"]), true
	} else {
		datas := GroupAutoReplyModel.Api_select_semi_byPercent(gid, rand)
		for _, data := range datas {
			if strings.Contains(text, Calc.Any2String(data["key"])) {
				return Calc.Any2String(data["value"]), true
			}
		}
	}
	return "", false
}
