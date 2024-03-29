package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/GroupFunction"
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/iapi"
	"main.go/app/bot/service"
	"main.go/config/app_default"
	"main.go/config/types"
	"main.go/tuuz/Redis"
)

func sign_daily() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + signDaily) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId
			GroupFunction.App_group_sign(self_id, group_id, user_id, message_id, gmr.GroupMember, gmr.GroupFunction)
		}
	}
}

func sign_lunpan() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + signLunpan) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId
			raw_message := gm.RawMessage
			GroupFunction.App_group_lunpan(self_id, group_id, user_id, message_id, raw_message, gmr.GroupMember, gmr.GroupFunction)
		}
	}
}

func check_score() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + checkScore) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId
			GroupFunction.App_check_balance(self_id, group_id, user_id, message_id, gmr.GroupMember, gmr.GroupFunction)
		}
	}
}

func rank_list() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + rankList) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId
			GroupFunction.App_check_rank(self_id, group_id, user_id, message_id, gmr.GroupMember, gmr.GroupFunction)
		}
	}
}

func word_limit() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + wordLimit) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId
			if gmr.GroupFunction["ban_retract"].(int64) == 1 {
				var rm iapi.RetractMessage
				rm.MessageId = message_id
				rm.SelfId = self_id
				rm.Time = 0
				ps.Publish_struct(types.RetractChannel, rm)
			}
			GroupFunction.App_ban_user(self_id, group_id, user_id, true, gmr.GroupFunction, app_default.Default_length_limit+"本群消息长度限制为："+Calc.Int642String(gmr.GroupFunction["word_limit"].(int64)))
		}
	}
}

func auto_reply() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + autoReply) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			//user_id := gm.UserId
			group_id := gm.GroupId
			//message_id := gm.MessageId
			raw_message := gm.RawMessage
			if str, ok := service.Serv_auto_reply(group_id, raw_message); ok {
				iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text(str), gmr.GroupFunction["auto_retract"].(int64) == 1)
			}
		}
	}
}

func greeting_when_at_me() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + atMe) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			//user_id := gm.UserId
			group_id := gm.GroupId
			//message_id := gm.MessageId
			raw_message := gm.RawMessage
			if _, ok := service.Serv_text_match_any(raw_message, []string{"[CQ:at,qq=" + Calc.Any2String(self_id) + "]"}); ok {
				iapi.Api.SendGroupMsg(self_id, group_id, MessageBuilder.IMessageBuilder{}.New().Text(app_default.Default_greetings), gmr.GroupFunction["auto_retract"].(int64) == 1)
			}
		}
	}
}

func daoju() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + 道具) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId
			raw_message := gm.RawMessage
			if str, ok := service.Serv_text_match(raw_message, []string{"道具"}); ok {
				GroupFunction.App_group_daoju(self_id, group_id, user_id, message_id, str, gmr.GroupMember, gmr.GroupFunction)
			}
		}
	}
}

func trade() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + 交易) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			message_id := gm.MessageId
			raw_message := gm.RawMessage
			if str, ok := service.Serv_text_match(raw_message, []string{"交易"}); ok {
				GroupFunction.App_trade_center(self_id, group_id, user_id, message_id, str, gmr.GroupMember, gmr.GroupFunction)
			}
		}
	}
}

func pal() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + palworld) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			//message_id := gm.MessageId
			raw_message := gm.RawMessage
			if str, ok := service.Serv_text_match(raw_message, []string{"交易"}); ok {
				GroupFunction.App_PalWorld(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(str), gmr.GroupFunction)
			}
		}
	}
}
