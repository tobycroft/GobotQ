package notice

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/iapi"
	group2 "main.go/app/bot/message/group"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupBlackListModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupKickModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/model/GroupMsgModel"
	"main.go/app/bot/model/LogErrorModel"
	"main.go/app/bot/model/LogRecvModel"
	"main.go/app/bot/service"
	"main.go/config/app_conf"
	"main.go/config/types"
	"main.go/tuuz"
	"main.go/tuuz/Jsong"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"net/netip"
	"time"
)

func Router() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageNotice) {
		var es EventStruct[Notice]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
		} else {
			em := es.Json
			bot := BotModel.Api_find(em.SelfId)
			if len(bot) < 1 {
				LogErrorModel.Api_insert("bot bot found", es.RemoteAddr)
				continue
			}
			ip := netip.MustParseAddrPort(es.RemoteAddr)
			if bot["allow_ip"] != ip.Addr().String() {
				LogErrorModel.Api_insert(fmt.Sprint("invalid ip address", bot["allow_ip"], ip.Addr().String()), em.SelfId)
				continue
			}
			notice_type := em.NoticeType
			sub_type := em.SubType
			group_id := em.GroupId
			self_id := em.SelfId

			groupfunction := GroupFunctionModel.Api_find(group_id)
			if len(groupfunction) < 1 {
				GroupFunctionModel.Api_insert(group_id)
				groupfunction = GroupFunctionModel.Api_find(group_id)
			}

			auto_retract := true
			if groupfunction["auto_retract"].(int64) == 0 {
				auto_retract = false
			}

			switch notice_type {
			//取消管理
			case "group_admin":
				var esg EventStruct[groupAdmin]
				err := sonic.UnmarshalString(c.Payload, &esg)
				if err != nil {
					LogErrorModel.Api_insert(err.Error(), c.Payload)
					break
				}
				ga := esg.Json
				user_id := ga.TargetId
				var group group2.RefreshGroupStruct
				group.GroupId = group_id
				group.SelfId = self_id
				group.UserId = user_id

				ps.Publish_struct(types.OperationRefreshGroup, group)
				switch sub_type {
				case "set":
					if user_id == self_id {
						if GroupMemberModel.Api_update_type(group_id, user_id, "admin") {
							go iapi.Api.Sendgroupmsg(self_id, group_id, "Acfur-On，已获取权限，可使用acfurhelp查看功能", auto_retract)
						} else {
							go iapi.Api.Sendgroupmsg(self_id, group_id, "Acfur-On，已获取权限，数据故障，请使用acfur刷新人数来更新信息", auto_retract)
						}
					} else {
						if GroupMemberModel.Api_update_type(group_id, user_id, "admin") {
							go iapi.Api.Sendgroupmsg(self_id, group_id, "恭喜上位"+service.Serv_at(user_id), auto_retract)
							GroupBlackListModel.Api_delete(group_id, user_id)
							GroupBanPermenentModel.Api_delete(group_id, user_id)
							Redis.Del("ban_" + Calc.Any2String(group_id) + "_" + Calc.Any2String(user_id))
						} else {
							go iapi.Api.Sendgroupmsg(self_id, group_id, "恭喜上位,但是权限变动失败", auto_retract)
						}
					}
					break

				case "unset":
					if user_id == self_id {
						if GroupMemberModel.Api_update_type(group_id, user_id, "member") {
							go iapi.Api.Sendgroupmsg(self_id, group_id, "Acfur-Off，权限已回收，将在2小时内退群", auto_retract)
						} else {
							go iapi.Api.Sendgroupmsg(self_id, group_id, "Acfur-Off，权限已回收，数据故障", auto_retract)
						}
					} else {
						if GroupMemberModel.Api_update_type(group_id, user_id, "member") {
							go iapi.Api.Sendgroupmsg(self_id, group_id, "管理员列表更新", auto_retract)
						} else {
							go iapi.Api.Sendgroupmsg(self_id, group_id, "管理员权限变动失败", auto_retract)
						}
					}
					break

				default:
					fmt.Println(em)
					break
				}
				break

			case "group_increase":
				var esg EventStruct[groupIncrease]
				err := sonic.UnmarshalString(c.Payload, &esg)
				if err != nil {
					LogErrorModel.Api_insert(err.Error(), c.Payload)
					break
				}
				ga := esg.Json
				user_id := ga.UserId
				if user_id == self_id {
					go Group.App_refreshmember(self_id, group_id)
				} else {
					if groupfunction["auto_hold"].(int64) == 1 {
						//如果禁言成功，就将这个人暂时加入永久小黑屋
						GroupBanPermenentModel.Api_insert(group_id, user_id, time.Now().Unix()+app_conf.Auto_ban_time-86400)
						num := Calc.Rand(1000, 9999)
						Redis.String_set("verify_"+Calc.Any2String(group_id)+"_"+Calc.Any2String(user_id), num, app_conf.Retract_time_duration+10*time.Second)
						Redis.String_set("ban_"+Calc.Any2String(group_id)+"_"+Calc.Any2String(user_id), num, 3600*time.Second)
						at := service.Serv_at(user_id)
						go iapi.Api.Sendgroupmsg(self_id, group_id, at+"请在120秒内在群内输入验证码数字：\n"+Calc.Any2String(num), true)
						go func(selfId, groupId, userId int64) {
							time.Sleep(120 * time.Second)
							ok, err := Redis.String_getBool("ban_" + Calc.Any2String(groupId) + "_" + Calc.Any2String(userId))
							if err != nil {
							} else {
								if ok {
									go iapi.Api.Sendgroupmsg(selfId, groupId, at+"看起来你没有完成活人验证，我现在将你加入永久小黑屋，但是你依然可以让其他管理员帮你解除", true)
									iapi.Api.SetGroupBan(selfId, groupId, userId, app_conf.Auto_ban_time)
								}
							}
						}(self_id, group_id, user_id)
						//Group.App_reverify(self_id, group_id, user_id, 0, "", nil, groupfunction)
					} else {
						//在没有启动自动验证模式的时候，使用正常欢迎流程
						if groupfunction["auto_welcome"].(int64) == 1 {
							if groupfunction["welcome_at"].(int64) == 1 {
								go iapi.Api.Sendgroupmsg(self_id, group_id, service.Serv_at(user_id)+Calc.Any2String(groupfunction["welcome_word"]), auto_retract)
							} else {
								go iapi.Api.Sendgroupmsg(self_id, group_id, Calc.Any2String(groupfunction["welcome_word"]), auto_retract)
							}
						} else {
							if groupfunction["join_alert"].(int64) == 1 {
								go iapi.Api.Sendgroupmsg(self_id, group_id, "成员+1", auto_retract)
							}
						}
					}

					if groupfunction["auto_card"].(int64) == 1 {
						comment, err := Redis.String_get("__request_comment__" + Calc.Any2String(group_id) + "_" + Calc.Any2String(user_id))
						if groupfunction["auto_card_insert"] == 1 && err == nil {
							go iapi.Api.Setgroupcard(self_id, group_id, user_id, comment)
						} else {
							go iapi.Api.Setgroupcard(self_id, group_id, user_id, groupfunction["auto_card_value"])
						}
					}

					go func(selfId, groupId, userId int64, autoRetract bool) {
						//将这个新加群的用户单条加入数据库
						member, err := iapi.Api.GetGroupMemberInfo(selfId, groupId, userId)
						if err != nil {

						} else {
							var mb GroupMemberModel.GroupMember
							mb.SelfId = selfId
							mb.UserId = userId
							mb.GroupId = groupId
							mb.Card = member.UserDisplayname
							mb.Title = member.Title
							mb.Level = member.Level
							mb.JoinTime = member.JoinTime
							mb.LastSentTime = member.LastSentTime
							mb.Nickname = member.Nickname
							mb.Role = member.Role
							if !GroupMemberModel.Api_insert(mb) {
								go iapi.Api.Sendgroupmsg(selfId, groupId, "群成员数据增加失败", autoRetract)
							}
						}
					}(self_id, group_id, user_id, auto_retract)
				}
				break

			case "group_decrease":
				var esg EventStruct[groupDecrease]
				err := sonic.UnmarshalString(c.Payload, &esg)
				if err != nil {
					LogErrorModel.Api_insert(err.Error(), c.Payload)
					break
				}
				ga := esg.Json
				operator_id := ga.OperatorId
				user_id := ga.UserId
				GroupMemberModel.Api_delete_byUid(self_id, group_id, user_id)
				switch sub_type {
				case "leave":
					if groupfunction["exit_to_black"].(int64) == 1 {
						GroupBlackListModel.Api_insert(group_id, user_id, operator_id)
						if groupfunction["exit_alert"].(int64) == 1 {
							go iapi.Api.Sendgroupmsg(self_id, group_id, Calc.Any2String(user_id)+"退群，已加入本群黑名单", auto_retract)
						}
					} else {
						if groupfunction["exit_alert"].(int64) == 1 {
							go iapi.Api.Sendgroupmsg(self_id, group_id, "成员-1", auto_retract)
						}
					}
					break

				case "kick":
					groupmsg := GroupMsgModel.Api_select(group_id, user_id, 10)
					last_msg := []string{}
					for _, data := range groupmsg {
						last_msg = append(last_msg, Calc.Any2String(data["text"]))
					}
					jsonmsg, _ := Jsong.Encode(last_msg)
					if groupfunction["kick_to_black"].(int64) == 1 {
						GroupBlackListModel.Api_insert(group_id, user_id, operator_id)
						if GroupKickModel.Api_insert(self_id, group_id, user_id, jsonmsg) {
							go iapi.Api.Sendgroupmsg(self_id, group_id, "群成员"+Calc.Any2String(user_id)+"T出报告已经生成，并已加入黑名单，请在APP中查看", auto_retract)
						} else {
							go iapi.Api.Sendgroupmsg(self_id, group_id, "群成员"+Calc.Any2String(user_id)+"T出报告生成失败，但已加入黑名单", auto_retract)
						}
					} else {
						if GroupKickModel.Api_insert(self_id, group_id, user_id, jsonmsg) {
							go iapi.Api.Sendgroupmsg(self_id, group_id, "群成员T出报告已经生成，请在APP中查看", auto_retract)
						} else {
							go iapi.Api.Sendgroupmsg(self_id, group_id, "群成员T出报告生成失败", auto_retract)
						}
					}
					break

				case "kick_me":
					GroupMemberModel.Api_delete_byGid(self_id, group_id)
					break

				case "group_recall":
					break

				default:
					fmt.Println("notice no route sub_type", em)
					break
				}
				break

			case "group_ban":
				switch sub_type {
				case "ban":
					var esg EventStruct[groupBan]
					err := sonic.UnmarshalString(c.Payload, &esg)
					if err != nil {
						LogErrorModel.Api_insert(err.Error(), c.Payload)
						break
					}
					ga := esg.Json
					user_id := ga.TargetId
					if ga.Duration >= 2505600 {
						if len(GroupBanPermenentModel.Api_find(group_id, user_id)) > 0 {

						} else {
							GroupBanPermenentModel.Api_insert(group_id, user_id, time.Now().Unix()+app_conf.Auto_ban_time-86400)
							go iapi.Api.Sendgroupmsg(self_id, group_id, service.Serv_at(user_id)+"你进入永久小黑屋，可联系群管解除", auto_retract)
						}
					}
					break

				case "lift_ban":
					var esg EventStruct[groupLiftBan]
					err := sonic.UnmarshalString(c.Payload, &esg)
					if err != nil {
						LogErrorModel.Api_insert(err.Error(), c.Payload)
						break
					}
					ga := esg.Json
					user_id := ga.TargetId
					if len(GroupBanPermenentModel.Api_find(group_id, user_id)) > 0 {
						GroupBanPermenentModel.Api_delete(group_id, user_id)
						go iapi.Api.Sendgroupmsg(self_id, group_id, service.Serv_at(user_id)+"你已经脱离永久小黑屋了", auto_retract)
					}
					break
				}
				break

			case "group_upload":
				fmt.Println("群上传", c.Payload)
				break

			//case "friend_add":
			//	//fmt.Println(em)
			//	break
			//
			//case "friend_recall":
			//	//fmt.Println(em)
			//	break
			//
			//case "group_recall":
			//	break
			//
			//case "notice":
			//	break

			default:
				fmt.Println("notice no route", em)
				LogRecvModel.Api_insert(c.Payload)
				break
			}
		}

	}

}
