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

func App_group_ban_word_set(bot, gid, uid int, text string, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	if len(text) > 2 {
		Type := text[0:1]
		new_str := text[1 : len(text)-1]
		if len(new_str) < 1 {
			api.Sendgroupmsg(bot, gid, "屏蔽词设定需要大于1位", true)
			return
		}
		switch Type {
		case "1":
			GroupBanWordModel.Api_insert(bot, gid, uid, new_str, 0, false, true, true, false)
			break

		case "2":
			GroupBanWordModel.Api_insert(bot, gid, uid, new_str, 0, true, false, true, false)
			break

		case "3":
			GroupBanWordModel.Api_insert(bot, gid, uid, new_str, 0, false, false, true, false)
			break

		case "-":
			data := GroupBanWordModel.Api_find(gid, new_str)
			if len(data) > 0 {
				if GroupBanWordModel.Api_delete(gid, new_str) {
					api.Sendgroupmsg(bot, gid, "屏蔽词已经删除", true)
				} else {
					api.Sendgroupmsg(bot, gid, "屏蔽词删除失败", true)
				}
			} else {
				api.Sendgroupmsg(bot, gid, "没有找到对应的屏蔽词，也许已经删除了呢？", true)
			}
			break

		case "+":
			break

		case "=":
			break

		default:
			break
		}
	} else {

	}
}
