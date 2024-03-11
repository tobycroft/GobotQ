package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/GroupFunction"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/service"
	"main.go/config/app_conf"
	"main.go/config/app_default"
	"main.go/config/types"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"regexp"
	"time"
)

func group_message_acfur_when_fully_matched() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupValid) {
		var es EventStruct[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Err(err)
		} else {
			gm := es.Json

			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId

			//the message in json format
			//message := gm.Message

			raw_message := gm.RawMessage

			text := raw_message
			reg := regexp.MustCompile("(?i)^acfur")
			active := reg.MatchString(text)
			new_text := reg.ReplaceAllString(text, "")
			groupmember := GroupMemberModel.Api_find(group_id, user_id)
			groupfunction := GroupFunctionModel.Api_find(group_id)
			if len(groupfunction) < 1 {
				GroupFunctionModel.Api_insert(group_id)
				groupfunction = GroupFunctionModel.Api_find(group_id)
			}
			if active {
				fmt.Println("acfuractive")
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
					go iapi.Api.SendGroupMsg(self_id, group_id, app_default.Default_greetings+time.Now().Format("2006-01-02 15:04:05"), true)
					break

				case "交易":
					GroupFunction.AutoMessage(self_id, group_id, user_id, app_default.Default_trade, groupfunction)
					break

				case "道具", "商店", "商城":
					GroupFunction.AutoMessage(self_id, group_id, user_id, app_default.Default_daoju, groupfunction)
					break

				case "轮盘":
					GroupFunction.AutoMessage(self_id, group_id, user_id, app_default.Default_lunpan_help, groupfunction)
					break

				case "help":
					GroupFunction.AutoMessage(self_id, group_id, user_id, app_default.Default_group_help, groupfunction)
					break

				case "app":
					GroupFunction.AutoMessage(self_id, group_id, user_id, app_default.Default_app_download_url, groupfunction)
					break

				case "设定":
					if !admin && !owner {
						service.Not_admin(self_id, group_id, user_id)
						break
					}
					GroupFunction.App_group_function_get_all(self_id, group_id, user_id, new_text, groupfunction)
					break

				case "刷新":
					GroupFunction.AutoMessage(self_id, group_id, user_id, "可以使用“刷新人数”以及“刷新群信息”来控制刷新", groupfunction)
					break

				case "权限", "查看权限":
					GroupFunction.AutoMessage(self_id, group_id, user_id, "我当前的权限为："+GroupFunction.BotPowerRefresh(group_id, self_id), groupfunction)
					break

				case "我的权限":
					GroupFunction.AutoMessage(self_id, group_id, user_id, "我当前的管理权限为："+Calc.Any2String(admin)+"\n群主权限为："+Calc.Any2String(owner), groupfunction)
					break

				case "查询群主", "查看群主", "呼叫群主":
					if !admin && !owner {
						service.Not_admin(self_id, group_id, user_id)
						break
					}
					owner_data := GroupMemberModel.Api_find_owner(self_id, group_id)
					if len(owner_data) > 0 {
						GroupFunction.AutoMessage(self_id, group_id, user_id, "本群群主为："+service.Serv_at(owner_data["user_id"]), groupfunction)
					} else {
						GroupFunction.AutoMessage(self_id, group_id, user_id, "本群未找到群主", groupfunction)
					}
					break

				case "随机数测试":
					rand1 := Calc.Rand(1, 100)
					rand2 := Calc.Rand(1, 100)
					GroupFunction.AutoMessage(self_id, group_id, user_id, "随机数1："+Calc.Any2String(rand1)+"\n随机数2："+Calc.Any2String(rand2), groupfunction)
					break

				case "刷新人数", "刷新群成员":
					if !admin && !owner {
						if len(groupmember) > 0 {
							service.Not_admin(self_id, group_id, user_id)
							break
						}
					}
					GroupFunction.App_refreshmember(self_id, group_id)
					GroupFunction.AutoMessage(self_id, group_id, user_id, "群用户已经刷新", groupfunction)
					break

				case "刷新群信息":
					if !admin && !owner {
						service.Not_admin(self_id, group_id, user_id)
						break
					}
					GroupFunction.App_refresh_groupinfo(self_id, group_id)
					GroupFunction.AutoMessage(self_id, group_id, user_id, "群信息刷新完成", groupfunction)
					break

				case "测试撤回":
					var rm iapi.RetractMessage
					rm.MessageId = message_id
					rm.SelfId = self_id
					rm.Time = 5 * time.Second
					if !admin {
						break
					}
					ps.Publish_struct(types.RetractChannel, rm)
					break

				case "测试T出测试":
					GroupFunction.App_kick_user(self_id, group_id, user_id, true, groupfunction, "测试")
					break

				case "测试禁言测试":
					GroupFunction.App_ban_user(self_id, group_id, user_id, true, groupfunction, "测试")
					break

				case "测试拼音":
					py, err := service.Serv_pinyin(new_text)
					if err != nil {

					} else {
						GroupFunction.AutoMessage(self_id, group_id, user_id, py, groupfunction)
					}
					break

				case "测试自动撤回":
					GroupFunction.AutoMessage(self_id, group_id, user_id, "自动撤回测试中……预计"+Calc.Any2String(app_conf.Retract_time_duration/time.Second+3)+"秒后撤回", groupfunction)
					break

				case "测试立即撤回":
					GroupFunction.AutoMessage(self_id, group_id, user_id, "自动撤回测试中……预计0秒后撤回", groupfunction)
					break

				case "屏蔽":
					if !admin && !owner {
						service.Not_admin(self_id, group_id, user_id)
						break
					}
					GroupFunction.AutoMessage(self_id, group_id, user_id, app_default.Default_str_ban_word, groupfunction)
					break

				case "屏蔽词":
					if !admin && !owner {
						service.Not_admin(self_id, group_id, user_id)
						break
					}
					GroupFunction.App_group_ban_word_list(self_id, group_id, user_id, new_text, 1, groupmember, groupfunction)
					break

				case "T出词":
					if !admin && !owner {
						service.Not_admin(self_id, group_id, user_id)
						break
					}
					GroupFunction.App_group_ban_word_list(self_id, group_id, user_id, new_text, 2, groupmember, groupfunction)
					break

				case "撤回词":
					if !admin && !owner {
						service.Not_admin(self_id, group_id, user_id)
						break
					}
					GroupFunction.App_group_ban_word_list(self_id, group_id, user_id, new_text, 3, groupmember, groupfunction)
					break

				case "清除小黑屋":
					if !admin && !owner {
						service.Not_admin(self_id, group_id, user_id)
						break
					}
					if GroupBanPermenentModel.Api_delete_byGroupId(group_id) {
						GroupFunction.AutoMessage(self_id, group_id, user_id, "小黑屋已经清除", groupfunction)
					} else {
						GroupFunction.AutoMessage(self_id, group_id, user_id, "小黑屋里面没有人啦~", groupfunction)
					}
					break

				case "查看人数", "查看人数上限":
					group_list_data := GroupListModel.Api_find(group_id)
					if len(group_list_data) > 0 {
						group_member_count := GroupMemberModel.Api_count_byGroupIdAndRole(group_id, nil)
						GroupFunction.AutoMessage(self_id, group_id, user_id, "本群人数上限为:"+Calc.Any2String(group_list_data["max_member_count"])+
							"\n当前人数为"+Calc.Any2String(group_member_count)+
							",\n如需清理请执行:acfur群人数清理", groupfunction)
					} else {
						GroupFunction.AutoMessage(self_id, group_id, user_id, "未找到本群，请使用acfur刷新群信息", groupfunction)
					}
					break

				case "群人数清理", "清理人数上限":
					if !owner && !admin {
						service.Not_owner(self_id, group_id, user_id)
						break
					}
					if Redis.CheckExists("__lock__group_id__" + Calc.Any2String(group_id)) {
						GroupFunction.AutoMessage(self_id, group_id, user_id, "执行中请稍等", groupfunction)
					} else {
						Redis.String_set("__lock__group_id__"+Calc.Any2String(group_id), 1, 60*time.Second)
						GroupFunction.App_drcrease_member(self_id, group_id, user_id, groupfunction, "")
					}
					break

				default:
					//if acfur triggered but fully match was failed, send to semi mode
					//ps.Publish(types.MessageGroupAcfur, c.Payload)
					break
				}
			} else {
				gmr := GroupMessageRedirect[GroupMessageStruct]{}
				gmr.GroupMember = groupmember
				gmr.GroupFunction = groupfunction
				gmr.Json = gm
				ps.Publish_struct(types.MessageGroupNormal, gmr)
			}

		}
	}
}
