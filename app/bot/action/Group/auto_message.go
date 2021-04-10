package Group

import "main.go/app/bot/api"

func AutoMessage(self_id, group_id, user_id interface{}, message string, groupfunction map[string]interface{}) {
	AutoRetract := false
	if groupfunction["auto_retract"].(int64) == 1 {
		AutoRetract = true
	}
	if groupfunction["all_send_private"].(int64) == 1 {
		api.Sendprivatemsg(self_id, user_id, group_id, message, AutoRetract)
	} else {
		api.Sendgroupmsg(self_id, group_id, message, AutoRetract)
	}
}
