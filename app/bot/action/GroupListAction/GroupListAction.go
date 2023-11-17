package GroupListAction

import (
	"main.go/app/bot/model/GroupListModel"
	"main.go/app/bot/redis/GroupListRedis"
)

func App_find_groupList(self_id, group_id any) GroupListModel.GroupList {
	data, err := GroupListRedis.Cac_find(self_id, group_id)
	if err != nil {
		data = GroupListModel.Api_find_struct(self_id, group_id)
		if data.GroupId != 0 {
			GroupListRedis.Cac_set(self_id, group_id, data)
			return data
		}
		return data
	} else {
		return data
	}
}
