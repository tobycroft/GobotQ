package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/bot/action/GroupFunction"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/redis/BanRepeatRedis"
	"main.go/app/bot/service"
	"main.go/config/types"
	"main.go/extend/Aigc"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
	"main.go/tuuz/Vali"
	"time"
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

			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId
			//message := gm.Message
			message := gm.RawMessage
			raw_message := gm.RawMessage

			groupfunction := gmr.GroupFunction

			var rm iapi.RetractMessage
			rm.MessageId = message_id
			rm.SelfId = self_id
			rm.Time = 0

			if err := Vali.Length(raw_message, -1, Calc.Any2Int64(groupfunction["word_limit"])); err != nil {
				ps.Publish_struct(types.MessageGroupAcfur+wordLimit, gmr)
			}
			if service.Serv_is_at_me(self_id, message) {
				ai_reply, err := Aigc.Aigc_gemini_text(message)
				if err != nil {
					fmt.Println(err)
					Log.Crrs(err, tuuz.FUNCTION_ALL())
				} else {
					iapi.Api.SendGroupMsg(self_id, group_id, Calc.Any2String(ai_reply), true)
				}
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
						GroupFunction.AutoMessage(selfId, groupId, userId, service.Serv_at(userId)+Calc.Any2String(groupFunction["repeat_time"])+"秒内请勿重复发送相同内容", groupFunction)
					}
				}
			}(self_id, group_id, user_id, groupfunction)

			go func(selfId, groupId, userId int64, groupFunction gorose.Data) {
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

						ps.Publish_struct(types.RetractChannel, rm)
						iapi.Api.SendGroupMsg(selfId, groupId, service.Serv_at(userId)+"验证成功"+str, true)
					} else {
						iapi.Api.SendGroupMsg(selfId, groupId, service.Serv_at(userId)+"你的输入不正确，需要输入："+Calc.Any2String(code), true)
					}
				}

				if Redis.CheckExists("ban_" + Calc.Any2String(groupId) + "_" + Calc.Any2String(userId)) {

					ps.Publish_struct(types.RetractChannel, rm)
					GroupFunction.AutoMessage(selfId, groupId, userId, service.Serv_at(userId)+"请尽快输入"+Calc.Any2String(code), groupFunction)
				} else if len(GroupBanPermenentModel.Api_find(groupId, userId)) > 0 {

					ps.Publish_struct(types.RetractChannel, rm)
					go iapi.Api.SendGroupMsg(self_id, group_id, "你现在处于永久小黑屋中，请让管理员使用acfur重新验证"+service.Serv_at(user_id)+"，来脱离当前状态", true)
				}
			}(self_id, group_id, user_id, groupfunction)
		}
	}
}
