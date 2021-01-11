package Group

import (
	"github.com/gohouse/gorose/v2"
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBanWordModel"
	"main.go/tuuz/Calc"
	"strings"
)

func App_group_ban_word_list(bot, gid, uid int, text string, Type int, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	kv := []gorose.Data{}
	switch Type {
	case 1:
		kv = GroupBanWordModel.Api_select_byKV(gid, "is_ban", true)
		break

	case 2:
		kv = GroupBanWordModel.Api_select_byKV(gid, "is_kick", true)
		break

	case 3:
		kv = GroupBanWordModel.Api_select_byKV(gid, "is_retract", true)
		break

	default:
		break
	}

	if len(kv) > 0 {
		word := []string{}
		for _, v := range kv {
			word = append(word, Calc.Any2String(v["word"]))
		}
		api.Sendgroupmsg(bot, gid, text+"列表为："+strings.Join(word, ","), true)
	} else {
		api.Sendgroupmsg(bot, gid, text+"列表为空，可以使用“acfur屏蔽”来查看设定方法", true)
	}
}
