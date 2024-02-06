package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/redis/BanRepeatRedis"
	"main.go/app/bot/service"
	"main.go/config/app_default"
	"main.go/config/types"
	"main.go/tuuz/Redis"
	"main.go/tuuz/Vali"
	"regexp"
	"time"
)

func group_message_acfur_semi_match() {
	go ban_group()
	go ban_url()
	go ban_wx()

	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur) {
		var es EventStruct[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := es.Json

			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId
			//message := gm.Message
			message := gm.RawMessage
			raw_message := gm.RawMessage

			text := message
			reg := regexp.MustCompile("(?i)^acfur")
			active := reg.MatchString(text)
			new_text := reg.ReplaceAllString(text, "")

			groupmember := GroupMemberModel.Api_find(group_id, user_id)
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

			groupfunction := GroupFunctionModel.Api_find(group_id)
			if len(groupfunction) < 1 {
				GroupFunctionModel.Api_insert(group_id)
				groupfunction = GroupFunctionModel.Api_find(group_id)
			}

			if active || service.Serv_is_at_me(self_id, message) {
				var rm iapi.RetractMessage
				rm.MessageId = message_id
				rm.SelfId = self_id
				rm.Time = 0

				gmr := GroupMessageRedirect[GroupMessageStruct]{}
				gmr.GroupMember = groupmember
				gmr.GroupFunction = groupfunction
				gmr.Json = gm
				if groupfunction["ban_group"].(int64) == 1 {
					if service.Serv_ban_group(raw_message) {
						if groupfunction["ban_retract"].(int64) == 1 {
							ps.Publish_struct(types.RetractChannel, rm)
						}
						fmt.Println(banGroup, self_id, group_id, user_id)
						ps.Publish_struct(types.MessageGroupAcfur+banGroup, gmr)
					}
				}

				if groupfunction["ban_url"].(int64) == 1 {
					if service.Serv_url_detect(raw_message) {
						if groupfunction["ban_retract"].(int64) == 1 {
							ps.Publish_struct(types.RetractChannel, rm)
						}
						fmt.Println(banUrl, self_id, group_id, user_id)
						ps.Publish_struct(types.MessageGroupAcfur+banUrl, gmr)
					}
				}

				if groupfunction["ban_wx"].(int64) == 1 {
					if service.Serv_ban_weixin(message) {
						if groupfunction["ban_retract"].(int64) == 1 {
							ps.Publish_struct(types.RetractChannel, rm)
						}
						fmt.Println(banWx, self_id, group_id, user_id)
						ps.Publish(types.MessageGroupAcfur+banWx, gmr)
					}
				}

				if groupfunction["ban_share"].(int64) == 1 {
					if service.Serv_ban_share(message) {
						if groupfunction["ban_retract"].(int64) == 1 {
							ps.Publish_struct(types.RetractChannel, rm)
						}
						fmt.Println(banShare, self_id, group_id, user_id)
						ps.Publish(types.MessageGroupAcfur+banShare, gmr)
					}
				}

				if _, ok := service.Serv_text_match(message, []string{"acfur屏蔽"}); ok {
					if !admin && !owner {
						if len(groupmember) > 0 {
							service.Not_admin(self_id, group_id, user_id)
						} else {
							Group.App_group_ban_word_set(self_id, group_id, user_id, message, message_id, groupmember, groupfunction)
							ps.Publish(types.MessageGroupAcfur+banWord, gmr)
						}
					}
				}

				if _, ok := service.Serv_text_match(message, []string{"acfur设定"}); ok {
					if !admin && !owner {
						if len(groupmember) > 0 {
							service.Not_admin(self_id, group_id, user_id)
						} else {
							Group.App_group_function_set(self_id, group_id, user_id, message, message_id, groupmember, groupfunction)
							ps.Publish(types.MessageGroupAcfur+setting, gmr)
						}
					}
				}

				if groupfunction["sign"].(int64) == 1 {
					if _, ok := service.Serv_text_match_all(message, []string{"签到"}); ok {
						ps.Publish(types.MessageGroupAcfur+sign, gmr)
						Group.App_group_sign(self_id, group_id, user_id, message_id, groupmember, groupfunction)
					}
				}

				if groupfunction["sign"].(int64) == 1 {
					if _, ok := service.Serv_text_match(message, []string{"轮盘"}); ok {
						ps.Publish(types.MessageGroupAcfur+轮盘, gmr)
						Group.App_group_lunpan(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
					}
				}

				if _, ok := service.Serv_text_match_all(message, []string{"积分查询", "查询积分", "威望查询", "查询威望", "钱包", "查询余额", "余额查询"}); ok {
					Group.App_check_balance(self_id, group_id, user_id, message_id, groupmember, groupfunction)
				}

				if _, ok := service.Serv_text_match_all(message, []string{"积分排行", "威望排行", "排行榜"}); ok {
					Group.App_check_rank(self_id, group_id, user_id, message_id, groupmember, groupfunction)
				}

				if err := Vali.Length(raw_message, -1, groupfunction["word_limit"].(int64)); err != nil {
					ps.Publish_struct(types.RetractChannel, rm)
					fmt.Println("长度限制", self_id, group_id, user_id)
					Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_length_limit+"本群消息长度限制为："+Calc.Int642String(groupfunction["word_limit"].(int64)))
				}

				if str, ok := service.Serv_auto_reply(group_id, raw_message); ok {
					iapi.Api.Sendgroupmsg(self_id, group_id, message, auto_retract)
				}
				if _, ok := service.Serv_text_match_any(message, []string{"[CQ:at,qq=" + Calc.Any2String(self_id) + "]"}); ok {
					iapi.Api.Sendgroupmsg(self_id, group_id, app_default.Default_welcome, true)
				}

				if str, ok := service.Serv_text_match(message, []string{"道具"}); ok {
					Group.App_group_daoju(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
				}

				if str, ok := service.Serv_text_match(message, []string{"交易"}); ok {
					Group.App_trade_center(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
				}

				if str, ok := service.Serv_text_match(message, []string{"acfur重新验证"}); ok {
					if !admin && !owner {
						if len(groupmember) > 0 {
							go service.Not_admin(self_id, group_id, user_id)
							break
						}
					}
					go Group.App_reverify(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
				}

				if str, ok := service.Serv_text_match(message, []string{"acfur死亡验证"}); ok {
					if !admin && !owner {
						if len(groupmember) > 0 {
							go service.Not_admin(self_id, group_id, user_id)
							break
						}
					}
					go Group.App_reverify_death(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
				}

				if str, ok := service.Serv_text_match(message, []string{"acfur活人验证"}); ok {
					if !admin && !owner {
						if len(groupmember) > 0 {
							go service.Not_admin(self_id, group_id, user_id)
							break
						}
					}
					go Group.App_reverify_force(self_id, group_id, user_id, message_id, message, groupmember, groupfunction)
				}
			}
		}
	}
}

func groupHandle_acfur_other(Type string, self_id, group_id, user_id, message_id int64, message, raw_message string, sender GroupSender, groupmember map[string]any, groupfunction map[string]any) {

	auto_retract := true
	if groupfunction["auto_retract"].(int64) == 0 {
		auto_retract = false
	}
	var ret iapi.Struct_Retract
	ret.MessageId = message_id
	ret.SelfId = self_id

	switch Type {

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
