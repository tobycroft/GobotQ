package cron

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/config/app_conf"
	"time"
)

func BanPermenentCheck() {
	datas := GroupBanPermenentModel.Api_select()
	for _, data := range datas {
		//设定下次禁言的时间为28天
		group_id := data["group_id"]
		user_id := data["user_id"]
		self_id := data["self_id"]
		gm := GroupMemberModel.Api_find(group_id, user_id)
		//如果这个用户不在群里面就不执行了
		if len(gm) > 0 {
			ok, _ := api.SetGroupBan(self_id, group_id, user_id, app_conf.Auto_ban_time)
			if ok {
				GroupBanPermenentModel.Api_update_nextTime(group_id, user_id, time.Now().Unix()+app_conf.Auto_ban_time)
			} else {
				//如果出现无法禁言的情况，默认就将这个人删除掉
				GroupBanPermenentModel.Api_delete(group_id, user_id)
			}
		}
	}
}
