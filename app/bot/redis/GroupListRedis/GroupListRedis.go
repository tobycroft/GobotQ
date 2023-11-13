package GroupListRedis

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/GroupListModel"
	"main.go/tuuz/Redis"
	"strings"
)

const GroupList = "GroupList"
const SelfId = ":SelfId:"
const GroupId = ":GroupId:"

func table(self_id, group_id any) string {
	str := strings.Builder{}
	str.WriteString(GroupList)
	str.WriteString(SelfId)
	if self_id != nil {
		str.WriteString(Calc.Any2String(self_id))
	} else {
		str.WriteString("*")
	}
	str.WriteString(GroupId)
	if group_id != nil {
		str.WriteString(Calc.Any2String(group_id))
	} else {
		str.WriteString("*")
	}
	return str.String()
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
