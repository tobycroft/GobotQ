package group

import (
	"fmt"
	"github.com/tobycroft/Calc"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/redis/BanRepeatRedis"
	"main.go/app/bot/service"
	"main.go/config/app_conf"
	"main.go/config/app_default"
	"main.go/tuuz/Redis"
	"sync"
	"time"
)

func groupHandle_acfur(self_id, group_id, user_id int64, message_id int64, new_text, message, raw_message string, sender GroupSender, groupmember map[string]any, groupfunction map[string]any) {
	admin := false
	owner := false
	if len(groupmember) > 0 {
		if groupmember["role"].(string) == "admin" {
			admin = true
		}
		if groupmember["role"].(string) == "owner" {
			admin = true
			owner = true
		}
	}
	switch new_text {

	case "":
		go iapi.Api.Sendgroupmsg(self_id, group_id, app_default.Default_welcome, true)
		break

	case "交易":
		Group.AutoMessage(self_id, group_id, user_id, app_default.Default_trade, groupfunction)
		break

	case "道具", "商店", "商城":
		Group.AutoMessage(self_id, group_id, user_id, app_default.Default_daoju, groupfunction)
		break

	case "轮盘":
		Group.AutoMessage(self_id, group_id, user_id, app_default.Default_lunpan_help, groupfunction)
		break

	case "help":
		Group.AutoMessage(self_id, group_id, user_id, app_default.Default_group_help, groupfunction)
		break

	case "app":
		Group.AutoMessage(self_id, group_id, user_id, app_default.Default_app_download_url, groupfunction)
		break

	case "设定":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			break
		}
		Group.App_group_function_get_all(self_id, group_id, user_id, new_text, groupfunction)
		break

	case "刷新":
		Group.AutoMessage(self_id, group_id, user_id, "可以使用“刷新人数”以及“刷新群信息”来控制刷新", groupfunction)
		break

	case "权限", "查看权限":
		Group.AutoMessage(self_id, group_id, user_id, "我当前的权限为："+Group.BotPowerRefresh(group_id, self_id), groupfunction)
		break

	case "我的权限":
		Group.AutoMessage(self_id, group_id, user_id, "我当前的管理权限为："+Calc.Any2String(admin)+"\n群主权限为："+Calc.Any2String(owner), groupfunction)
		break

	case "查询群主", "查看群主", "呼叫群主":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			break
		}
		owner_data := GroupMemberModel.Api_find_owner(self_id, group_id)
		if len(owner_data) > 0 {
			Group.AutoMessage(self_id, group_id, user_id, "本群群主为："+service.Serv_at(owner_data["user_id"]), groupfunction)
		} else {
			Group.AutoMessage(self_id, group_id, user_id, "本群未找到群主", groupfunction)
		}
		break

	case "随机数测试":
		rand1 := Calc.Rand(1, 100)
		rand2 := Calc.Rand(1, 100)
		Group.AutoMessage(self_id, group_id, user_id, "随机数1："+Calc.Any2String(rand1)+"\n随机数2："+Calc.Any2String(rand2), groupfunction)
		break

	case "刷新人数", "刷新群成员":
		if !admin && !owner {
			if len(groupmember) > 0 {
				service.Not_admin(self_id, group_id, user_id)
				break
			}
		}
		Group.App_refreshmember(self_id, group_id)
		Group.AutoMessage(self_id, group_id, user_id, "群用户已经刷新", groupfunction)
		break

	case "刷新群信息":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			break
		}
		Group.App_refresh_groupinfo(self_id, group_id)
		Group.AutoMessage(self_id, group_id, user_id, "群信息刷新完成", groupfunction)
		break

	case "测试撤回":
		var ret iapi.Struct_Retract
		ret.MessageId = message_id
		ret.SelfId = self_id
		if !admin {
			break
		}
		go func(ret iapi.Struct_Retract) {
			iapi.Retract_instant <- ret
		}(ret)
		break

	case "测试T出测试":
		Group.App_kick_user(self_id, group_id, user_id, true, groupfunction, "测试")
		break

	case "测试禁言测试":
		Group.App_ban_user(self_id, group_id, user_id, true, groupfunction, "测试")
		break

	case "测试拼音":
		py, err := service.Serv_pinyin(new_text)
		if err != nil {

		} else {
			Group.AutoMessage(self_id, group_id, user_id, py, groupfunction)
		}
		break

	case "测试自动撤回":
		Group.AutoMessage(self_id, group_id, user_id, "自动撤回测试中……预计"+Calc.Int2String(app_conf.Retract_time_second+3)+"秒后撤回", groupfunction)
		break

	case "测试立即撤回":
		Group.AutoMessage(self_id, group_id, user_id, "自动撤回测试中……预计0秒后撤回", groupfunction)
		break

	case "屏蔽":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			break
		}
		Group.AutoMessage(self_id, group_id, user_id, app_default.Default_str_ban_word, groupfunction)
		break

	case "屏蔽词":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			break
		}
		Group.App_group_ban_word_list(self_id, group_id, user_id, new_text, 1, groupmember, groupfunction)
		break

	case "T出词":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			break
		}
		Group.App_group_ban_word_list(self_id, group_id, user_id, new_text, 2, groupmember, groupfunction)
		break

	case "撤回词":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			break
		}
		Group.App_group_ban_word_list(self_id, group_id, user_id, new_text, 3, groupmember, groupfunction)
		break

	case "清除小黑屋":
		if !admin && !owner {
			service.Not_admin(self_id, group_id, user_id)
			break
		}
		if GroupBanPermenentModel.Api_delete_byGroupId(group_id) {
			Group.AutoMessage(self_id, group_id, user_id, "小黑屋已经清除", groupfunction)
		} else {
			Group.AutoMessage(self_id, group_id, user_id, "小黑屋里面没有人啦~", groupfunction)
		}
		break

	case "查看人数", "查看人数上限":
		group_list_data := GroupListModel.Api_find(group_id)
		if len(group_list_data) > 0 {
			group_member_count := GroupMemberModel.Api_count_byGroupIdAndRole(group_id, nil)
			Group.AutoMessage(self_id, group_id, user_id, "本群人数上限为:"+Calc.Any2String(group_list_data["max_member_count"])+
				"\n当前人数为"+Calc.Any2String(group_member_count)+
				",\n如需清理请执行:acfur群人数清理", groupfunction)
		} else {
			Group.AutoMessage(self_id, group_id, user_id, "未找到本群，请使用acfur刷新群信息", groupfunction)
		}
		break

	case "群人数清理", "清理人数上限":
		if !owner && !admin {
			service.Not_owner(self_id, group_id, user_id)
			break
		}
		if Redis.CheckExists("__lock__group_id__" + Calc.Any2String(group_id)) {
			Group.AutoMessage(self_id, group_id, user_id, "执行中请稍等", groupfunction)
		} else {
			Redis.String_set("__lock__group_id__"+Calc.Any2String(group_id), 1, 60*time.Second)
			Group.App_drcrease_member(self_id, group_id, user_id, groupfunction, "")
		}
		break

	default:
		groupHandle_acfur_middle(self_id, group_id, user_id, message_id, message, raw_message, sender, groupmember, groupfunction)
		break
	}
}

var group_function_type = []string{"unknow", "ban_group", "url_detect", "ban_weixin", "ban_share", "ban_word", "setting",
	"sign", "轮盘", "威望查询", "威望排行", "长度限制", "自动回复", "atme", "道具", "交易", "重新验证", "死亡验证", "活人验证"}

func groupHandle_acfur_middle(self_id, group_id, user_id, message_id int64, message, raw_message string, sender GroupSender, groupmember map[string]any, groupfunction map[string]any) {
	function := make([]bool, len(group_function_type)+1, len(group_function_type)+1)
	new_text := make([]string, len(group_function_type)+1, len(group_function_type)+1)
	var wg sync.WaitGroup
	wg.Add(len(group_function_type))
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		new_text[idx] = message
	}(0, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		ok := service.Serv_ban_group(raw_message)
		new_text[idx] = raw_message
		function[idx] = ok
	}(1, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		ok := service.Serv_url_detect(raw_message)
		new_text[idx] = raw_message
		function[idx] = ok
	}(2, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		ok := service.Serv_ban_weixin(message)
		new_text[idx] = message
		function[idx] = ok
	}(3, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		ok := service.Serv_ban_share(message)
		new_text[idx] = message
		function[idx] = ok
	}(4, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"acfur屏蔽"})
		new_text[idx] = str
		function[idx] = ok
	}(5, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"acfur设定"})
		new_text[idx] = str
		function[idx] = ok
	}(6, &wg)
	//签到(直接)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match_all(message, []string{"签到"})
		new_text[idx] = str
		function[idx] = ok
	}(7, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"轮盘"})
		new_text[idx] = str
		function[idx] = ok
	}(8, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match_all(message, []string{"积分查询", "查询积分", "威望查询", "查询威望", "钱包", "查询余额", "余额查询"})
		new_text[idx] = str
		function[idx] = ok
	}(9, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		_, ok := service.Serv_text_match_all(message, []string{"积分排行", "威望排行", "排行榜"})
		new_text[idx] = ""
		function[idx] = ok
	}(10, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		if int64(len(raw_message)) > groupfunction["word_limit"].(int64) {
			new_text[idx] = raw_message
			function[idx] = true
		}
	}(11, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_auto_reply(group_id, raw_message)
		new_text[idx] = str
		function[idx] = ok
	}(12, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		_, ok := service.Serv_text_match_any(message, []string{"[CQ:at,qq=" + Calc.Any2String(self_id) + "]"})
		new_text[idx] = ""
		function[idx] = ok
	}(13, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"道具"})
		new_text[idx] = str
		function[idx] = ok
	}(14, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"交易"})
		new_text[idx] = str
		function[idx] = ok
	}(15, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"acfur重新验证"})
		new_text[idx] = str
		function[idx] = ok
	}(16, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"acfur死亡验证"})
		new_text[idx] = str
		function[idx] = ok
	}(17, &wg)
	go func(idx int, wg *sync.WaitGroup) {
		defer wg.Done()
		str, ok := service.Serv_text_match(message, []string{"acfur活人验证"})
		new_text[idx] = str
		function[idx] = ok
	}(18, &wg)
	wg.Wait()
	function_route := 0
	for i := range function {
		if function[i] == true {
			function_route = i
			break
		}
	}
	groupHandle_acfur_other(group_function_type[function_route], self_id, group_id, user_id, message_id, new_text[function_route], raw_message, sender, groupmember, groupfunction)
}

func groupHandle_acfur_other(Type string, self_id, group_id, user_id, message_id int64, message, raw_message string, sender GroupSender, groupmember map[string]any, groupfunction map[string]any) {
	admin := false
	owner := false
	if len(groupmember) > 0 {
		if groupmember["role"].(string) == "admin" {
			admin = true
		}
		if groupmember["role"].(string) == "owner" {
			admin = true
			owner = true
		}
	}
	auto_retract := true
	if groupfunction["auto_retract"].(int64) == 0 {
		auto_retract = false
	}
	var ret iapi.Struct_Retract
	ret.MessageId = message_id
	ret.SelfId = self_id

	switch Type {

	case "重新验证":
		if !admin && !owner {
			if len(groupmember) > 0 {
				go service.Not_admin(self_id, group_id, user_id)
				break
			}
		}
		go Group.App_reverify(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
		break

	case "活人验证":
		if !admin && !owner {
			if len(groupmember) > 0 {
				go service.Not_admin(self_id, group_id, user_id)
				break
			}
		}
		go Group.App_reverify_force(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
		break

	case "死亡验证":
		if !admin && !owner {
			if len(groupmember) > 0 {
				go service.Not_admin(self_id, group_id, user_id)
				break
			}
		}
		go Group.App_reverify_death(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
		break

	case "交易":
		go Group.App_trade_center(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
		break

	case "道具":
		go Group.App_group_daoju(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
		break

	case "atme":
		go iapi.Api.Sendgroupmsg(self_id, group_id, app_default.Default_welcome, true)
		break

	case "sign":
		if groupfunction["sign"].(int64) == 1 {
			go Group.App_group_sign(self_id, group_id, user_id, message_id, groupmember, groupfunction)
		}
		break

	case "轮盘":
		if groupfunction["sign"].(int64) == 1 {
			go Group.App_group_lunpan(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
		}
		break

	case "setting":
		if !admin && !owner {
			if len(groupmember) > 0 {
				go service.Not_admin(self_id, group_id, user_id)
				break
			}
		}
		go Group.App_group_function_set(self_id, group_id, user_id, message, message_id, groupmember, groupfunction)
		break

	case "ban_word":
		if !admin && !owner {
			if len(groupmember) > 0 {
				go service.Not_admin(self_id, group_id, user_id)
				break
			}
		}
		go Group.App_group_ban_word_set(self_id, group_id, user_id, message, message_id, groupmember, groupfunction)
		break

	case "url_detect":
		if groupfunction["ban_url"].(int64) == 1 {
			if groupfunction["ban_retract"].(int64) == 1 {
				go func(ret iapi.Struct_Retract) {
					iapi.Retract_instant <- ret
				}(ret)
			}
			fmt.Println("url_detect", self_id, group_id, user_id)
			go Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_url)
		}
		break

	case "ban_group":
		if groupfunction["ban_group"].(int64) == 1 {
			if groupfunction["ban_retract"].(int64) == 1 {
				go func(ret iapi.Struct_Retract) {
					iapi.Retract_instant <- ret
				}(ret)
			}
			fmt.Println("ban_group", self_id, group_id, user_id)
			go Group.App_kick_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_group)
		}
		break

	case "ban_weixin":
		if groupfunction["ban_wx"].(int64) == 1 {
			if groupfunction["ban_retract"].(int64) == 1 {
				go func(ret iapi.Struct_Retract) {
					iapi.Retract_instant <- ret
				}(ret)
			}
			fmt.Println("ban_weixin", self_id, group_id, user_id)
			go Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_weixin)
		}
		break

	case "ban_share":
		if groupfunction["ban_share"].(int64) == 1 {
			if groupfunction["ban_retract"].(int64) == 1 {
				go func(ret iapi.Struct_Retract) {
					iapi.Retract_instant <- ret
				}(ret)
			}
			fmt.Println("ban_share", self_id, group_id, user_id)
			go Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_share)
		}
		break

	case "威望查询":
		go Group.App_check_balance(self_id, group_id, user_id, message_id, groupmember, groupfunction)
		break

	case "威望排行":
		go Group.App_check_rank(self_id, group_id, user_id, message_id, groupmember, groupfunction)
		break

	case "长度限制":
		go func(ret iapi.Struct_Retract) {
			iapi.Retract_instant <- ret
		}(ret)
		fmt.Println("长度限制", self_id, group_id, user_id)
		go Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction,
			app_default.Default_length_limit+"本群消息长度限制为："+Calc.Int642String(groupfunction["word_limit"].(int64)))
		break

	case "自动回复":
		go iapi.Api.Sendgroupmsg(self_id, group_id, message, auto_retract)
		break

	default:

		go func(selfId, groupId, userId any, groupFunction gorose.Data) {
			if groupFunction["ban_repeat"].(int64) == 1 {
				num, err := BanRepeatRedis.BanRepeatRedis{}.Table(userId, raw_message).Cac_find()
				if err != nil {
					num = 0
				}
				BanRepeatRedis.BanRepeatRedis{}.Table(userId, raw_message).Cac_set(num+1, time.Duration(groupFunction["repeat_time"].(int64))*time.Second)
				if num > groupFunction["repeat_count"].(int64) {
					Group.App_ban_user(selfId, groupId, userId, auto_retract, groupFunction, "请不要在"+Calc.Any2String(groupFunction["repeat_time"])+"秒内重复发送相同内容")
				} else if int64(num)+1 > groupFunction["repeat_count"].(int64) {
					Group.AutoMessage(selfId, groupId, userId, service.Serv_at(userId)+Calc.Any2String(groupFunction["repeat_time"])+"秒内请勿重复发送相同内容", groupFunction)
				}
			}
		}(self_id, group_id, user_id, groupfunction)

		go func(selfId, groupId, userId any, groupFunction gorose.Data) {
			//验证程序
			code, err := Redis.String_get("verify_" + Calc.Any2String(groupId) + "_" + Calc.Any2String(userId))
			if err != nil {
			} else {
				if code == message {
					str := ""
					Redis.Del("ban_" + Calc.Any2String(groupId) + "_" + Calc.Any2String(userId))
					Redis.Del("verify_" + Calc.Any2String(groupId) + "_" + Calc.Any2String(userId))
					if len(GroupBanPermenentModel.Api_find(groupId, userId)) > 0 {
						GroupBanPermenentModel.Api_delete(groupId, userId)
						str += "\r\n永久小黑屋记录已移除"
					}
					if groupFunction["auto_welcome"] == 1 {
						str = "\r\n" + Calc.Any2String(groupFunction["welcome_word"])
					}
					go func(ret iapi.Struct_Retract) {
						iapi.Retract_instant <- ret
					}(ret)
					iapi.Api.Sendgroupmsg(selfId, groupId, service.Serv_at(userId)+"验证成功"+str, true)
				} else {
					iapi.Api.Sendgroupmsg(selfId, groupId, service.Serv_at(userId)+"你的输入不正确，需要输入："+Calc.Any2String(code), true)
				}
			}

			if Redis.CheckExists("ban_" + Calc.Any2String(groupId) + "_" + Calc.Any2String(userId)) {
				go func(ret iapi.Struct_Retract) {
					iapi.Retract_instant <- ret
				}(ret)
				Group.AutoMessage(selfId, groupId, userId, service.Serv_at(userId)+"请尽快输入"+Calc.Any2String(code), groupFunction)
			} else if len(GroupBanPermenentModel.Api_find(groupId, userId)) > 0 {
				go func(ret iapi.Struct_Retract) {
					iapi.Retract_instant <- ret
				}(ret)
				go iapi.Post{}.Sendgroupmsg(self_id, group_id, "你现在处于永久小黑屋中，请让管理员使用acfur重新验证"+service.Serv_at(user_id)+"，来脱离当前状态", true)
			}
		}(self_id, group_id, user_id, groupfunction)

		break
	}
}