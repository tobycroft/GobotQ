package Group

import (
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/GroupMemberModel"
)

func BotPower(group_id, self_id any) string {
	member_bot := GroupMemberModel.Api_find(group_id, self_id)
	if len(member_bot) > 0 {
		return member_bot["role"].(string)
	}
	return "member"
}

func BotPowerRefresh(group_id, self_id any) string {
	gm, err := iapi.Api{}.GetGroupMemberInfo(self_id, group_id, self_id)
	if err != nil {

	} else {
		GroupMemberModel.Api_update_type(group_id, self_id, gm.Role)
	}
	member_bot := GroupMemberModel.Api_find(group_id, self_id)
	if len(member_bot) > 0 {
		return member_bot["role"].(string)
	}
	return "member"
}
