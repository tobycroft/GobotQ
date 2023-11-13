package GroupListRedis

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/GroupListModel"
	"main.go/tuuz/Redis"
)

const GroupList = "GroupList"
const SelfId = ":SelfId:"
const GroupId = ":GroupId:"

func table(self_id, group_id any) string {
	return GroupList + SelfId + Calc.Any2String(self_id) + GroupId + Calc.Any2String(group_id)
}
func Cac_set[T GroupListModel.GroupList](self_id, group_id any, data T) error {
	return Redis.Hash_set_struct(table(self_id, group_id), data)
}

func Cac_find[T GroupListModel.GroupList](self_id, group_id any) (T, error) {
	var data T
	err := Redis.Hash_get_struct(table(self_id, group_id), &data)
	return data, err
}

func Cac_del(self_id, group_id any) error {
	return Redis.Hash_delete(table(self_id, group_id))
}
