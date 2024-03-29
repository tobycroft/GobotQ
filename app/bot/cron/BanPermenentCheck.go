package cron

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/GroupBanPermenentModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/config/app_conf"
	"time"
)

func BanPermenentCheck() {
	for {
		ban_permenent_check()
		time.Sleep(1 * time.Hour)
	}
}

func ban_permenent_check() {
	datas := GroupBanPermenentModel.Api_select()
	for _, data := range datas {
		//设定下次禁言的时间为28天
		group_id := Calc.Any2Int64(data["group_id"])
		user_id := Calc.Any2Int64(data["user_id"])
		self_id := Calc.Any2Int64(data["self_id"])
		gm := GroupMemberModel.Api_find(group_id, user_id)
		//如果这个用户不在群里面就不执行了
		if len(gm) > 0 {
			ok, _ := iapi.Api.SetGroupBan(self_id, group_id, user_id, app_conf.Auto_ban_time)
			if ok {
				//如果禁言成功了就把这个人的禁言时间延长即可
				GroupBanPermenentModel.Api_update_nextTime(group_id, user_id, time.Now().Unix()+app_conf.Auto_ban_time-86400)
			} else {
				//如果出现无法禁言的情况，默认就将这个人删除掉
				GroupBanPermenentModel.Api_delete(group_id, user_id)
			}
		}
	}
}
