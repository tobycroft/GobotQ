package GroupMemberRedis

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/tuuz/Redis"
)

const GroupMember = "GroupMember"
const SelfId = ":SelfId:"
const GroupId = ":GroupId:"
const UserId = ":UserId:"

func table(self_id, user_id, group_id any) string {
	return GroupMember + SelfId + Calc.Any2String(self_id) + UserId + Calc.Any2String(user_id) + GroupId + Calc.Any2String(group_id)
}
func Cac_set[T GroupMemberModel.GroupMember](self_id, user_id, group_id any, data T) error {
	return Redis.Hash_set_struct(table(self_id, user_id, group_id), data)
}

func Cac_find[T GroupMemberModel.GroupMember](self_id, user_id, group_id any) (T, error) {
	var data T
	err := Redis.Hash_get_struct(table(self_id, user_id, group_id), &data)
	return data, err
}

func Cac_del(self_id, user_id, group_id any) error {
	return Redis.Hash_delete(table(self_id, user_id, group_id))
}
