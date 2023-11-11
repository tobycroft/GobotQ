package Group

import (
	"main.go/app/bot/iapi/apipost"
)

func AutoMessage(self_id, group_id, user_id any, message string, groupfunction map[string]any) {
	AutoRetract := false
	if groupfunction["auto_retract"].(int64) == 1 {
		AutoRetract = true
	}
	if groupfunction["all_send_private"].(int64) == 1 {
		go apipost.Api{}.Sendprivatemsg(self_id, user_id, group_id, message, AutoRetract)
	} else {
		go apipost.Api{}.Sendgroupmsg(self_id, group_id, message, AutoRetract)
	}
}
