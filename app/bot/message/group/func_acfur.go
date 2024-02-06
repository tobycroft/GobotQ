package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/iapi"
	"main.go/config/app_default"
	"main.go/config/types"
	"main.go/tuuz/Redis"
)

func ban_group() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + banGroup) {
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
			Group.App_kick_user(self_id, group_id, user_id, true, gmr.GroupFunction, app_default.Default_ban_group)
		}
	}
}

func ban_url() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + banUrl) {
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
			Group.App_ban_user(self_id, group_id, user_id, true, gmr.GroupFunction, app_default.Default_ban_url)
		}
	}
}

func ban_wx() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + banWx) {
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
			Group.App_ban_user(self_id, group_id, user_id, true, gmr.GroupFunction, app_default.Default_ban_weixin)
		}
	}
}

func ban_share() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + banShare) {
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
			Group.App_ban_user(self_id, group_id, user_id, true, gmr.GroupFunction, app_default.Default_ban_share)
		}
	}
}

func ban_word() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + banWord) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			raw_message := gm.RawMessage
			Group.App_group_ban_word_set(self_id, group_id, user_id, raw_message)
		}
	}
}

func settings() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + setting) {
		var gmr GroupMessageRedirect[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &gmr)
		if err != nil {
			fmt.Println(err)
		} else {
			gm := gmr.Json
			self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			raw_message := gm.RawMessage
			Group.App_group_function_set(self_id, group_id, user_id, raw_message, gmr.GroupFunction)
		}
	}
}

func signs() {
	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupAcfur + sign) {
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
			Group.App_group_sign(self_id, group_id, user_id, message_id, gmr.GroupMember, gmr.GroupFunction)
		}
	}
}
