package FriendListAction

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/app/bot/model/FriendListModel"
	"main.go/app/bot/redis/FriendListRedis"
)

func App_find_friendList(user_id any) (gorose.Data, error) {
	e, err := FriendListRedis.Cac_find[gorose.Data](user_id)
	if err != nil {
		data := FriendListModel.Api_find(user_id)
		if len(data) > 0 {
			FriendListRedis.Cac_set(user_id, data)
			return FriendListRedis.Cac_find[gorose.Data](user_id)
		} else {
			return nil, err
		}
	} else {
		return e, err
	}
}
