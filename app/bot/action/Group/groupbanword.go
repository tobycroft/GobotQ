package Group

import (
	"github.com/tobycroft/Calc"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/GroupBanWordModel"

	"strings"
)

func App_group_ban_word_list(self_id, group_id, user_id int64, message string, Type int, groupmember map[string]any, groupfunction map[string]any) {
	kv := []gorose.Data{}
	switch Type {
	case 1:
		kv = GroupBanWordModel.Api_select_byKV(group_id, "is_ban", true)
		break

	case 2:
		kv = GroupBanWordModel.Api_select_byKV(group_id, "is_kick", true)
		break

	case 3:
		kv = GroupBanWordModel.Api_select_byKV(group_id, "is_retract", true)
		break

	default:
		break
	}

	if len(kv) > 0 {
		word := []string{}
		for _, v := range kv {
			word = append(word, Calc.Any2String(v["word"]))
		}
		go iapi.Api.Sendgroupmsg(self_id, group_id, message+"列表为："+strings.Join(word, ","), true)
	} else {
		go iapi.Api.Sendgroupmsg(self_id, group_id, message+"列表为空，可以使用“acfur屏蔽”来查看设定方法", true)
	}
}

func App_group_ban_word_set(self_id, group_id, user_id int64, message string) {
	if len(message) > 2 {
		text_slice := strings.Split(message, "")
		Type := text_slice[0]
		new_str := ""
		for i := 1; i < len(text_slice); i++ {
			new_str += text_slice[i]
		}
		if len(new_str) < 1 {
			go iapi.Api.Sendgroupmsg(self_id, group_id, "屏蔽词设定需要大于1位", true)
			return
		}
		switch Type {
		case "1":
			data := GroupBanWordModel.Api_find(group_id, new_str)
			if len(data) > 0 {
				if GroupBanWordModel.Api_update(group_id, new_str, data["is_kick"], true, true) {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "修改成功", true)
				} else {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "修改失败", true)
				}
			} else {
				if GroupBanWordModel.Api_insert(group_id, user_id, new_str, 0, false, true, true, false) {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "添加成功", true)
				} else {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "添加失败", true)
				}
			}
			break

		case "2":
			data := GroupBanWordModel.Api_find(group_id, new_str)
			if len(data) > 0 {
				if GroupBanWordModel.Api_update(group_id, new_str, true, data["is_ban"], true) {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "修改成功", true)
				} else {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "修改失败", true)
				}
			} else {
				if GroupBanWordModel.Api_insert(group_id, user_id, new_str, 0, true, false, true, false) {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "添加成功", true)
				} else {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "添加失败", true)
				}
			}
			break

		case "3":
			data := GroupBanWordModel.Api_find(group_id, new_str)
			if len(data) > 0 {
				if GroupBanWordModel.Api_update(group_id, new_str, data["is_kick"], data["is_ban"], true) {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "修改成功", true)
				} else {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "修改失败", true)
				}
			} else {
				if GroupBanWordModel.Api_insert(group_id, user_id, new_str, 0, false, false, true, false) {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "添加成功", true)
				} else {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "添加失败", true)
				}
			}
			break

		case "-":
			data := GroupBanWordModel.Api_find(group_id, new_str)
			if len(data) > 0 {
				if GroupBanWordModel.Api_delete(group_id, new_str) {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "屏蔽词已经删除", true)
				} else {
					go iapi.Api.Sendgroupmsg(self_id, group_id, "屏蔽词删除失败", true)
				}
			} else {
				go iapi.Api.Sendgroupmsg(self_id, group_id, "没有找到对应的屏蔽词，也许已经删除了呢？", true)
			}
			break

		case "+":
			strs := strings.Split(new_str, "#")
			if len(strs) < 2 {
				go iapi.Api.Sendgroupmsg(self_id, group_id, "设定错误，你可以使用：acfur屏蔽+屏蔽词#处罚，例如acfur屏蔽+触发词#T出撤回", true)
			} else {
				data := GroupBanWordModel.Api_find(group_id, strs[0])
				if len(data) > 0 {
					if GroupBanWordModel.Api_update(group_id, strs[0], strings.Contains(strs[1], "T出"), strings.Contains(strs[1], "屏蔽"), strings.Contains(strs[1], "撤回")) {
						go iapi.Api.Sendgroupmsg(self_id, group_id, "屏蔽词更新成功", true)
					} else {
						go iapi.Api.Sendgroupmsg(self_id, group_id, "屏蔽词更新失败", true)
					}
				} else {
					if GroupBanWordModel.Api_insert(group_id, user_id, strs[0], 0, strings.Contains(strs[1], "T出"),
						strings.Contains(strs[1], "屏蔽"), strings.Contains(strs[1], "撤回"), false) {
						go iapi.Api.Sendgroupmsg(self_id, group_id, "屏蔽词添加成功，新增："+strs[0], true)
					} else {
						go iapi.Api.Sendgroupmsg(self_id, group_id, "屏蔽词"+strs[0]+"添加失败", true)
					}
				}

			}
			break

		default:
			go iapi.Api.Sendgroupmsg(self_id, group_id, "如需添加，请使用acfur屏蔽(+)1/2/3/+/-/=(+)屏蔽词即可，如需其他帮助，请发送acfur屏蔽来查看添加方法", true)
			break
		}
	} else {
		go iapi.Api.Sendgroupmsg(self_id, group_id, "屏蔽词命令有误，请发送acfur屏蔽查看如何使用屏蔽词", true)
	}
}
