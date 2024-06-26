package notice

import (
	"errors"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/GroupFunction"
	"main.go/app/bot/action/MessageBuilder"
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
	"main.go/config/app_conf"
	"main.go/config/types"
	"main.go/extend/TTS"
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
				var esg EventStruct[GroupAdmin]
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
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("Acfur-On，已获取权限，可使用acfurhelp查看功能"), auto_retract)
						} else {
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("Acfur-On，已获取权限，数据故障，请使用acfur刷新人数来更新信息"), auto_retract)
						}
					} else {
						if GroupMemberModel.Api_update_type(group_id, user_id, "admin") {
							msg := MessageBuilder.IMessageBuilder{}.New().Text("恭喜上位").At(user_id)
							go iapi.Api.SendGroupMsg(self_id, group_id, msg, auto_retract)
							GroupBlackListModel.Api_delete(group_id, user_id)
							GroupBanPermenentModel.Api_delete(group_id, user_id)
							Redis.Del("ban_" + Calc.Any2String(group_id) + "_" + Calc.Any2String(user_id))
						} else {
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("恭喜上位,但是权限变动失败"), auto_retract)
						}
					}
					break

				case "unset":
					if user_id == self_id {
						if GroupMemberModel.Api_update_type(group_id, user_id, "member") {
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("Acfur-Off，权限已回收，将在2小时内退群"), auto_retract)
						} else {
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("Acfur-Off，权限已回收，数据故障"), auto_retract)
						}
					} else {
						if GroupMemberModel.Api_update_type(group_id, user_id, "member") {
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("管理员列表更新"), auto_retract)
						} else {
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("管理员权限变动失败"), auto_retract)
						}
					}
					break

				default:
					fmt.Println(em)
					break
				}
				break

			case "group_increase":
				var esg EventStruct[GroupIncrease]
				err := sonic.UnmarshalString(c.Payload, &esg)
				if err != nil {
					LogErrorModel.Api_insert(err.Error(), c.Payload)
					break
				}
				ga := esg.Json
				user_id := ga.UserId
				if user_id == 0 {
					LogErrorModel.Api_insert(errors.New("user_id=0"), c.Payload)
					break
				}
				if user_id == self_id {
					go GroupFunction.App_refreshmember(self_id, group_id)
				} else {
					iapi.Api.GetGroupMemberInfo(self_id, group_id, user_id)
					time.Sleep(2 * time.Second)

					if groupfunction["auto_hold"].(int64) == 1 {
						//如果禁言成功，就将这个人暂时加入永久小黑屋
						GroupBanPermenentModel.Api_insert(group_id, user_id, time.Now().Unix()+app_conf.Auto_ban_time-86400)
						num := Calc.Rand(1000, 9999)
						Redis.String_set("verify_"+Calc.Any2String(group_id)+"_"+Calc.Any2String(user_id), num, app_conf.Retract_time_duration+10*time.Second)
						Redis.String_set("ban_"+Calc.Any2String(group_id)+"_"+Calc.Any2String(user_id), num, 3600*time.Second)
						msg := MessageBuilder.IMessageBuilder{}.New().At(user_id)
						go iapi.Api.SendGroupMsg(self_id, group_id, msg.Text("请在120秒内在群内输入验证码数字：\n"+Calc.Any2String(num)), true)
						go func(selfId, groupId, userId int64) {
							time.Sleep(120 * time.Second)
							ok, err := Redis.String_getBool("ban_" + Calc.Any2String(groupId) + "_" + Calc.Any2String(userId))
							if err != nil {
							} else {
								if ok {
									go iapi.Api.SendGroupMsg(selfId, groupId, msg.Text("看起来你没有完成活人验证，我现在将你加入永久小黑屋，但是你依然可以让其他管理员帮你解除"), true)
									iapi.Api.SetGroupBan(selfId, groupId, userId, app_conf.Auto_ban_time)
								}
							}
						}(self_id, group_id, user_id)
						//Group.App_reverify(self_id, group_id, user_id, 0, "", nil, groupfunction)
					} else {
						//在没有启动自动验证模式的时候，使用正常欢迎流程
						if Calc.Any2Int64(groupfunction["auto_welcome"]) == 1 {
							if Calc.Any2Int64(groupfunction["welcome_voice"]) == 1 {
								usr := GroupMemberModel.Api_find(group_id, user_id)
								name := ""
								if len(usr) > 0 {
									name += Calc.Any2String(usr["nickname"])
								}
								msg := MessageBuilder.IMessageBuilder{}.New()
								audio, err := TTS.Audio{}.New().Huihui(Calc.Any2String(name + "，" + Calc.Any2String(groupfunction["welcome_word"])))
								if err != nil {
									//msg.Text(err.Error())
									Log.Crrs(err, tuuz.FUNCTION_ALL())
								} else {
									msg.Record(audio.AudioUrl)
									iapi.Api.SendGroupMsg(self_id, group_id, msg, auto_retract)
								}
							} else {
								msg := MessageBuilder.IMessageBuilder{}.New().Text(Calc.Any2String(groupfunction["welcome_word"]))
								if groupfunction["welcome_at"].(int64) == 1 {
									iapi.Api.SendGroupMsg(self_id, group_id, msg.At(user_id), auto_retract)
								} else {
									iapi.Api.SendGroupMsg(self_id, group_id, msg, auto_retract)
								}
							}
						} else {
							if groupfunction["join_alert"].(int64) == 1 {
								iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("成员+1"), auto_retract)
							}
						}
					}

					if groupfunction["auto_card"].(int64) == 1 {
						comment, err := Redis.String_get("__request_comment__" + Calc.Any2String(group_id) + "_" + Calc.Any2String(user_id))
						if groupfunction["auto_card_insert"] == 1 && err == nil {
							go iapi.Api.SetGroupCard(self_id, group_id, user_id, comment)
						} else {
							go iapi.Api.SetGroupCard(self_id, group_id, user_id, groupfunction["auto_card_value"])
						}
					}

				}
				break

			case "group_decrease":
				var esg EventStruct[GroupDecrease]
				err := sonic.UnmarshalString(c.Payload, &esg)
				if err != nil {
					LogErrorModel.Api_insert(err.Error(), c.Payload)
					break
				}
				ga := esg.Json
				operator_id := ga.OperatorId
				user_id := ga.UserId
				if user_id == 0 {
					LogErrorModel.Api_insert(errors.New("user_id=0"), c.Payload)
					break
				}
				GroupMemberModel.Api_delete_byUid(self_id, group_id, user_id)
				switch sub_type {
				case "leave":
					if groupfunction["exit_to_black"].(int64) == 1 {
						GroupBlackListModel.Api_insert(group_id, user_id, operator_id)
						if groupfunction["exit_alert"].(int64) == 1 {
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text(Calc.Any2String(user_id)+"退群，已加入本群黑名单"), auto_retract)
						}
					} else {
						if groupfunction["exit_alert"].(int64) == 1 {
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("成员-1"), auto_retract)
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
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("群成员"+Calc.Any2String(user_id)+"T出报告已经生成，并已加入黑名单，请在APP中查看"), auto_retract)
						} else {
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("群成员"+Calc.Any2String(user_id)+"T出报告生成失败，但已加入黑名单"), auto_retract)
						}
					} else {
						if GroupKickModel.Api_insert(self_id, group_id, user_id, jsonmsg) {
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("群成员T出报告已经生成，请在APP中查看"), auto_retract)
						} else {
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("群成员T出报告生成失败"), auto_retract)
						}
					}
					break

				case "kick_me":
					GroupMemberModel.Api_delete_byGid(self_id, group_id)
					break

				default:
					fmt.Println("notice no route sub_type", em)
					break
				}
				break

			case "group_ban":
				switch sub_type {
				case "ban":
					var esg EventStruct[GroupBan]
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
							go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().At(user_id).Text("你进入永久小黑屋，可联系群管解除"), auto_retract)
						}
					}
					break

				case "lift_ban":
					var esg EventStruct[GroupLiftBan]
					err := sonic.UnmarshalString(c.Payload, &esg)
					if err != nil {
						LogErrorModel.Api_insert(err.Error(), c.Payload)
						break
					}
					ga := esg.Json
					user_id := ga.TargetId
					if len(GroupBanPermenentModel.Api_find(group_id, user_id)) > 0 {
						GroupBanPermenentModel.Api_delete(group_id, user_id)
						go iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().At(user_id).Text("你已经脱离永久小黑屋了"), auto_retract)
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

			case "group_recall":
				var esg EventStruct[groupRecallMessage]
				err := sonic.UnmarshalString(c.Payload, &esg)
				if err != nil {
					LogErrorModel.Api_insert(err.Error(), c.Payload)
					break
				}
				ga := esg.Json
				fmt.Println("群成员撤回消息", ga)
				break

			//case "notice":
			//	break

			case "notify":
				switch sub_type {
				case "poke":
					break

				}
				break

			default:
				fmt.Println("notice no route", em)
				LogRecvModel.Api_insert(c.Payload)
				break
			}
		}

	}

}

type groupRecallMessage struct {
	OperatorId int    `json:"operator_id"`
	UserId     int    `json:"user_id"`
	TipText    string `json:"tip_text"`
	NoticeType string `json:"notice_type"`
	GroupId    int    `json:"group_id"`
	Time       int    `json:"time"`
	SelfId     int    `json:"self_id"`
	MessageId  int    `json:"message_id"`
	PostType   string `json:"post_type"`
}
