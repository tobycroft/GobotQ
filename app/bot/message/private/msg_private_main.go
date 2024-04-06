package private

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/PrivateAutoReplyModel"
	"main.go/config/types"
	"main.go/extend/Aigc"
	"main.go/extend/STT"
	"main.go/extend/TTS"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

func message_main_handler() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessagePrivateValid) {
		var es EventStruct[PrivateMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Errs(err, tuuz.FUNCTION_ALL())
		} else {
			pm := es.Json
			selfId := pm.SelfId
			user_id := pm.UserId
			group_id := int64(0)
			message := pm.Message
			use_voice := false
			normal_text := strings.Builder{}
			for _, msg := range message {
				switch msg.Type {
				case "at":
					break

				case "text":
					normal_text.WriteString(Calc.Any2String(msg.Data["text"]))
					break

				case "record":
					go func() {
						time.Sleep(500 * time.Millisecond)
						iapi.Api.GetRecord(selfId, Calc.Any2String(msg.Data["file"]), "mp3")
					}()
					fmt.Println("语音解析:", Calc.Any2String(msg.Data["file"]))
					select {
					case <-time.NewTicker(10 * time.Second).C:
						iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("语音解析超时"), false)
						break

					case c := <-Redis.PubSub{}.Subscribe(types.GetFile + Calc.Any2String(msg.Data["file"])):
						fmt.Println("接收语音:", Calc.Any2String(msg.Data["file"]))
						if c.Payload == "fail" {
							iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text("b64解码失败"), false)
							break
						}
						str, err := STT.Audio{}.New().SpeechBase64ToText(c.Payload)
						if err != nil {
							iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text(err.Error()), false)
							break
						}
						fmt.Println("语音解析", str)
						use_voice = true
						normal_text.WriteString(str)
						break

					}
				}
			}

			text := normal_text.String()

			reg := regexp.MustCompile("(?i)^acfur")
			active := reg.MatchString(text)
			if !active {
				//在未激活acfur的情况下应该对原始内容进行还原
				re, ok := private_default_reply(text)
				if ok {
					if use_voice {
						if use_voice {
							rec, err := TTS.Audio{}.New().Huihui(re)
							if err == nil {
								iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.New().Record(rec.AudioUrl), false)
								continue
							}
						}
					}
					continue
				}
				auto_reply := PrivateAutoReplyModel.Api_find_byKey(text)
				if len(auto_reply) > 0 {
					if auto_reply["value"] != nil {
						if use_voice {
							rec, err := TTS.Audio{}.New().Huihui(auto_reply["value"].(string))
							if err == nil {
								iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.New().Record(rec.AudioUrl), false)
								continue
							}
						}
						iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text(auto_reply["value"].(string)), false)
						continue
					}
				} else {
					re, ok := private_auto_reply(selfId, text)
					if ok {
						if use_voice {
							if use_voice {
								rec, err := TTS.Audio{}.New().Huihui(re)
								if err == nil {
									iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.New().Record(rec.AudioUrl), false)
									continue
								}
							}
						}
						continue
					}
				}
				if utf8.RuneCountInString(text) > 2 {
					ai_reply, err := Aigc.Aigc_bing_text(text)
					if err != nil {
						Log.Crrs(err, tuuz.FUNCTION_ALL())
						iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text(err.Error()), false)
						continue
					}
					if use_voice {
						rec, err := TTS.Audio{}.New().Huihui(ai_reply.Data)
						if err == nil {
							iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.New().Record(rec.AudioUrl), false)
							continue
						}
					}
					iapi.Api.SendPrivateMsg(selfId, user_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text(ai_reply.Echo), false)
					continue
				}
			}
		}
	}
}
