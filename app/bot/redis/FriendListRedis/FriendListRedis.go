package FriendListRedis

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/FriendListModel"
	"main.go/tuuz/Redis"
)

const FriendInfo = "FriendInfo"
const SelfId = ":SelfId:"
const UserId = ":UserId:"

func table(self_id, user_id any) string {
	return FriendInfo + SelfId + Calc.Any2String(self_id) + UserId + Calc.Any2String(user_id)
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
