package Group

import (
	"fmt"
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
		text_slice := strings.Split(text, "")
		Type := text_slice[0]
		fmt.Println(text_slice)
		new_str := ""
		for i := 1; i < len(text_slice); i++ {
			new_str += text_slice[i]
		}
		if len(new_str) < 1 {
			api.Sendgroupmsg(bot, gid, "屏蔽词设定需要大于1位", true)
			return
		}
		switch Type {
		case "1":
			data := GroupBanWordModel.Api_find(gid, new_str)
			if len(data) > 0 {
				if GroupBanWordModel.Api_update(gid, new_str, data["is_kick"], true, true) {
					api.Sendgroupmsg(bot, gid, "修改成功", true)
				} else {
					api.Sendgroupmsg(bot, gid, "修改失败", true)
				}
			} else {
				if GroupBanWordModel.Api_insert(bot, gid, uid, new_str, 0, false, true, true, false) {
					api.Sendgroupmsg(bot, gid, "添加成功", true)
				} else {
					api.Sendgroupmsg(bot, gid, "添加失败", true)
				}
			}
			break

		case "2":
			data := GroupBanWordModel.Api_find(gid, new_str)
			if len(data) > 0 {
				if GroupBanWordModel.Api_update(gid, new_str, true, data["is_ban"], true) {
					api.Sendgroupmsg(bot, gid, "修改成功", true)
				} else {
					api.Sendgroupmsg(bot, gid, "修改失败", true)
				}
			} else {
				if GroupBanWordModel.Api_insert(bot, gid, uid, new_str, 0, true, false, true, false) {
					api.Sendgroupmsg(bot, gid, "添加成功", true)
				} else {
					api.Sendgroupmsg(bot, gid, "添加失败", true)
				}
			}
			break

		case "3":
			data := GroupBanWordModel.Api_find(gid, new_str)
			if len(data) > 0 {
				if GroupBanWordModel.Api_update(gid, new_str, data["is_kick"], data["is_ban"], true) {
					api.Sendgroupmsg(bot, gid, "修改成功", true)
				} else {
					api.Sendgroupmsg(bot, gid, "修改失败", true)
				}
			} else {
				if GroupBanWordModel.Api_insert(bot, gid, uid, new_str, 0, false, false, true, false) {
					api.Sendgroupmsg(bot, gid, "添加成功", true)
				} else {
					api.Sendgroupmsg(bot, gid, "添加失败", true)
				}
			}
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
			strs := strings.Split(new_str, "#")
			if len(strs) < 2 {
				api.Sendgroupmsg(bot, gid, "设定错误，你可以使用：acfur屏蔽+屏蔽词#处罚，例如acfur屏蔽+触发词#T出撤回", true)
			} else {
				data := GroupBanWordModel.Api_find(gid, new_str)
				if len(data) > 0 {
					if GroupBanWordModel.Api_update(gid, strs[0], strings.Contains(strs[1], "T出"), strings.Contains(strs[1], "屏蔽"), strings.Contains(strs[1], "撤回")) {
						api.Sendgroupmsg(bot, gid, "屏蔽词被成功", true)
					} else {
						api.Sendgroupmsg(bot, gid, "屏蔽词修改失败", true)
					}
				} else {
					if GroupBanWordModel.Api_insert(bot, gid, uid, strs[0], 0, strings.Contains(strs[1], "T出"),
						strings.Contains(strs[1], "屏蔽"), strings.Contains(strs[1], "撤回"), false) {
						api.Sendgroupmsg(bot, gid, "屏蔽词添加成功，新增："+strs[0], true)
					} else {
						api.Sendgroupmsg(bot, gid, "屏蔽词"+strs[0]+"添加失败", true)
					}
				}

			}
			break

		default:
			api.Sendgroupmsg(bot, gid, "如需添加，请使用acfur屏蔽(+)1/2/3/+/-/=(+)屏蔽词即可，如需其他帮助，请发送acfur屏蔽来查看添加方法", true)
			break
		}
	} else {
		api.Sendgroupmsg(bot, gid, "屏蔽词命令有误，请发送acfur屏蔽查看如何使用屏蔽词", true)
	}
}
