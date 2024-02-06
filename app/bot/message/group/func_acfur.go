package group

import (
	"main.go/app/bot/action/Group"
	"main.go/config/app_default"
)

func ban_group() {

	Group.App_kick_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_group)

}

func ban_url() {
	Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_url)

}

func ban_wx() {
	go Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_weixin)

}

func ban_share() {
	go Group.App_ban_user(self_id, group_id, user_id, auto_retract, groupfunction, app_default.Default_ban_share)

}
