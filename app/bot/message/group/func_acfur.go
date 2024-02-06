package group

import (
	"fmt"
	"main.go/app/bot/action/Group"
	"main.go/app/bot/iapi"
	"main.go/config/app_default"
)

func ban_group() {
	if groupfunction["ban_group"].(int64) == 1 {
		if groupfunction["ban_retract"].(int64) == 1 {
			go func(ret iapi.Struct_Retract) {
				iapi.Retract_instant <- ret
			}(ret)
		}
		fmt.Println("ban_group", self_id, group_id, user_id)
		go Group.App_kick_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_group)
	}
}
