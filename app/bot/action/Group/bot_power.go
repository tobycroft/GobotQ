package Group

import "main.go/app/bot/model/GroupMemberModel"

func BotPower(group_id, self_id interface{}) string {
	member_bot := GroupMemberModel.Api_find(group_id, self_id)
	if len(member_bot) > 0 {
		return member_bot["role"].(string)
	}
	return "member"
}
