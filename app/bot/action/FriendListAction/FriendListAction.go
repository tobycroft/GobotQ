package FriendListAction

import (
	"main.go/app/bot/model/FriendListModel"
	"main.go/app/bot/redis/FriendListRedis"
)

func App_find_friendList[T FriendListModel.FriendList](self_id, user_id any) (T, error) {
	e, err := FriendListRedis.Cac_find[T](self_id, user_id)
	if err != nil {
		data := FriendListModel.Api_find_struct[T](user_id)
		if *new(T) != data {
			FriendListRedis.Cac_set(self_id, user_id, data)
			return data, nil
		}
		return data, err
	} else {
		return e, err
	}
}
