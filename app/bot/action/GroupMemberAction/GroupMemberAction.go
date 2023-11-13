package GroupMemberAction

import (
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/redis/GroupMemberRedis"
)

func App_find_groupMember[T GroupMemberModel.GroupMember](self_id, user_id, group_id any) (T, error) {
	data, err := GroupMemberRedis.Cac_find[T](self_id, user_id, group_id)
	if err != nil {
		e := GroupMemberModel.Api_find_struct(self_id, user_id, group_id)
		if e != *new(GroupMemberModel.GroupMember) {
			GroupMemberRedis.Cac_set(e.SelfId, e.UserId, e.GroupId, e)
			return data, nil
		}
		return data, err
	} else {
		return data, err
	}
}
