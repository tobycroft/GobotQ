package group

import (
	"main.go/app/bot/action/Group"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/service"
	"main.go/config/app_default"
	"regexp"
	"time"
)

func GroupHandle(self_id, group_id, user_id, message_id int64, message, raw_message string, sender GroupSender) {
	text := message
	reg := regexp.MustCompile("(?i)^acfur")
	active := reg.MatchString(text)
	new_text := reg.ReplaceAllString(text, "")
	groupmember := GroupMemberModel.Api_find(group_id, user_id)
	groupfunction := GroupFunctionModel.Api_find(group_id)
	if len(groupfunction) < 1 {
		GroupFunctionModel.Api_insert(group_id)
		groupfunction = GroupFunctionModel.Api_find(group_id)
	}
	botinfo := BotModel.Api_find(self_id)

	if active || service.Serv_is_at_me(self_id, message) {
		if botinfo["end_date"].(time.Time).Before(time.Now()) {
			Group.AutoMessage(self_id, group_id, user_id, app_default.Default_over_time, groupfunction)
			return
		}
		go groupHandle_acfur(self_id, group_id, user_id, message_id, new_text, message, raw_message, sender, groupmember, groupfunction)
	} else {
		if botinfo["end_date"].(time.Time).Before(time.Now()) {
			return
		}
		//在未激活acfur的情况下应该对原始内容进行还原
		go groupHandle_acfur_middle(self_id, group_id, user_id, message_id, message, raw_message, sender, groupmember, groupfunction)
	}
}
