package FriendListRedis

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/FriendListModel"
	"main.go/tuuz/Redis"
	"strings"
)

const FriendList = "FriendList"
const SelfId = ":SelfId:"
const UserId = ":UserId:"

func table(self_id, user_id any) string {
	str := strings.Builder{}
	str.WriteString(FriendList)
	str.WriteString(SelfId)
	if self_id != nil {
		str.WriteString(Calc.Any2String(self_id))
	} else {
		str.WriteString("*")
	}
	str.WriteString(UserId)
	if user_id != nil {
		str.WriteString(Calc.Any2String(user_id))
	} else {
		str.WriteString("*")
	}
	return str.String()
}

func Cac_set[T FriendListModel.FriendList](self_id, user_id any, data T) error {
	return Redis.Hash_set_struct(table(self_id, user_id), data)
}

func Cac_find[T FriendListModel.FriendList](self_id, user_id any) (T, error) {
	var data T
	err := Redis.Hash_get_struct(table(self_id, user_id), &data)
	return data, err
}

func Cac_del(self_id, user_id any) error {
	return Redis.Hash_delete(table(self_id, user_id))
}
