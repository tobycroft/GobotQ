package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/bot/action/GroupFunction"
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/redis/BanRepeatRedis"
	"main.go/config/types"
	"main.go/extend/Aigc"
	"main.go/extend/STT"
	"main.go/extend/TTS"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"main.go/tuuz/Vali"
	"strings"
	"time"
	"unicode/utf8"
)

// group_message_normal 处理标注消息
func group_message_normal() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupNormal) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			Log.Err(err)
		} else {
			gm := gmr.Json
			groupfunction := gmr.GroupFunction

			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId
			message := gm.Message
			raw_message := gm.RawMessage

			normal_text := strings.Builder{}
			is_at_me := false
			use_voice := false
			for _, msg := range message {
				switch msg.Type {
				case "at":
					if Calc.Any2String(msg.Data["qq"]) == Calc.Any2String(self_id) {
						is_at_me = true
					}
					break

				case "text":
					normal_text.WriteString(msg.Data["text"])
					break

				case "record":
					go func() {
						time.Sleep(500 * time.Millisecond)
						iapi.Api.GetRecord(self_id, msg.Data["file"], "mp3")
					}()

					fmt.Println("语音解析:", msg.Data["file"])
					select {
					case <-time.NewTicker(10 * time.Second).C:
						iapi.Api.SendPrivateMsg(self_id, user_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("语音解析超时"), true)
						break

					case c := <-Redis.PubSub{}.Subscribe(types.GetFile + msg.Data["file"]):
						fmt.Println("接收语音:", msg.Data["file"])
						if c.Payload == "fail" {
							GroupFunction.AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text("b64解码失败"), groupfunction)
							break
						}
						str, err := STT.Audio{}.New().SpeechBase64ToText(c.Payload)
						if err != nil {
							GroupFunction.AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(err.Error()), groupfunction)
							break
						}
						fmt.Println("语音解析", str)
						use_voice = true
						normal_text.WriteString(str)
						break

					}

				}
			}

			var rm iapi.RetractMessage
			rm.MessageId = message_id
			rm.SelfId = self_id
			rm.Time = 0

			text := normal_text.String()
			if err := Vali.Length(raw_message, -1, Calc.Any2Int64(groupfunction["word_limit"])); err != nil {
				ps.Publish_struct(types.MessageGroupAcfur+wordLimit, gmr)
			}
			if is_at_me {
				if utf8.RuneCountInString(text) > 4 {
					ai_reply, err := Aigc.Aigc_bing_text(text)
					if err != nil {
						fmt.Println(err)
						Log.Crrs(err, tuuz.FUNCTION_ALL())
						GroupFunction.AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(err.Error()), groupfunction)
					} else {
						GroupFunction.AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().At(user_id).Text(ai_reply.Echo), groupfunction)
					}
				}
			} else if use_voice {
				b := make(chan bool, 0)
				go func() {
					run_time := 0
					for {
						select {
						case <-time.NewTicker(10 * time.Second).C:
							run_time += 1
							if run_time <= 1 {
								usr := GroupMemberModel.Api_find(group_id, user_id)
								name := ""
								if len(usr) > 0 {
									name += Calc.Any2String(usr["nickname"])
								}
								rec, err := TTS.Audio{}.New().Huihui("请稍等一下" + TTS.Audio{}.CleanText(name) + "，我正在生成回答，可能需要一些时间")
								if err != nil {
									GroupFunction.AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(err.Error()), groupfunction)
								} else {
									iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Record(rec.AudioUrl), true)
								}
							}
							break

						case <-b:
							return
						}
					}
				}()

				go func() {
					ai_reply, err := Aigc.Aigc_bing_text(text)
					b <- true
					if err != nil {
						fmt.Println(err)
						Log.Crrs(err, tuuz.FUNCTION_ALL())
						GroupFunction.AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(err.Error()), groupfunction)
					} else {
						rec, err := TTS.Audio{}.New().Huihui(ai_reply.Data)
						if err != nil {
							GroupFunction.AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(err.Error()), groupfunction)
						} else {
							GroupFunction.AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Record(rec.AudioUrl), groupfunction)
						}
					}
				}()

			}
			go func(selfId, groupId, userId int64, groupFunction gorose.Data) {
				if Calc.Any2Int64(groupFunction["ban_repeat"]) == 1 {
					num, err := BanRepeatRedis.BanRepeatRedis{}.Table(userId, raw_message).Cac_find()
					if err != nil {
						num = 0
					}
					BanRepeatRedis.BanRepeatRedis{}.Table(userId, raw_message).Cac_set(num+1, time.Duration(groupFunction["repeat_time"].(int64))*time.Second)
					if num > groupFunction["repeat_count"].(int64) {
						GroupFunction.App_ban_user(selfId, groupId, userId, groupfunction["auto_retract"].(int64) == 1, groupFunction, "请不要在"+Calc.Any2String(groupFunction["repeat_time"])+"秒内重复发送相同内容")
					} else if int64(num)+1 > groupFunction["repeat_count"].(int64) {
						GroupFunction.AutoMessage(selfId, groupId, userId, MessageBuilder.IMessageBuilder{}.New().At(userId).Text(Calc.Any2String(groupFunction["repeat_time"])+"秒内请勿重复发送相同内容"), groupFunction)
					}
				}
			}(self_id, group_id, user_id, groupfunction)

			go func(selfId, groupId, userId int64, groupFunction gorose.Data) {
				//验证程序
				code, err := Redis.String_get("verify_" + Calc.Any2String(groupId) + "_" + Calc.Any2String(userId))
				if err != nil {
				} else {
					if code == text {
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

						ps.Publish_struct(types.RetractChannel, rm)
						iapi.Api.SendGroupMsg(selfId, groupId, MessageBuilder.IMessageBuilder{}.New().At(userId).Text("验证成功"+str), true)
					} else {
						iapi.Api.SendGroupMsg(selfId, groupId, MessageBuilder.IMessageBuilder{}.New().At(userId).Text("你的输入不正确，需要输入："+Calc.Any2String(code)), true)
					}
				}

				if Redis.CheckExists("ban_" + Calc.Any2String(groupId) + "_" + Calc.Any2String(userId)) {
					ps.Publish_struct(types.RetractChannel, rm)
					GroupFunction.AutoMessage(selfId, groupId, userId, MessageBuilder.IMessageBuilder{}.New().At(userId).Text("请尽快输入"+Calc.Any2String(code)), groupFunction)
				} else if len(GroupBanPermenentModel.Api_find(groupId, userId)) > 0 {
					ps.Publish_struct(types.RetractChannel, rm)
					msg := MessageBuilder.IMessageBuilder{}.New().New().Text("你现在处于永久小黑屋中，请让管理员使用acfur重新验证").At(user_id).Text("，来脱离当前状态")
					go iapi.Api.SendGroupMsg(self_id, group_id, msg, true)
				}
			}(self_id, group_id, user_id, groupfunction)
		}
	}
}
