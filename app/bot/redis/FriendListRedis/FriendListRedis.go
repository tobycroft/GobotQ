package FriendListRedis

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/model/FriendListModel"
	"main.go/tuuz/Redis"
)

const FriendInfo = "FriendInfo:"

func table(user_id any) string {
	return FriendInfo + Calc.Any2String(user_id)
}
func Cac_set[T FriendListModel.FriendList](user_id any, data T) error {
	return Redis.Hash_set_struct(table(user_id), data)
}

func Cac_find[T FriendListModel.FriendList](user_id any) (T, error) {
	var data T
	err := Redis.Hash_get_struct(table(user_id), &data)
	return data, err
}

func Cac_delete(user_id any) error {
	return Redis.Hash_delete(table(user_id))
}
