package GroupListAction

import (
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/redis/GroupListRedis"
)

func App_find_groupList[T GroupListModel.GroupList](self_id, group_id any) (T, error) {
	data, err := GroupListRedis.Cac_find[T](self_id, group_id)
	if err != nil {
		data = GroupListModel.Api_find_struct[T](self_id, group_id)
		if data != nil {
			GroupListRedis.Cac_set(self_id, group_id, data)
			return data, nil
		}
		return data, err
	} else {
		return data, err
	}
}
