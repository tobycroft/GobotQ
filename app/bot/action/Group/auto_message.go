package Group

import "main.go/app/bot/api"

func AutoMessage(self_id, group_id, user_id interface{}, message string, AutoRetract bool, groupfunction map[string]interface{}) {
	if groupfunction["all_send_private"].(int64) == 1 {
		api.Sendprivatemsg(self_id, user_id, group_id, message, AutoRetract)
	} else {
		api.Sendgroupmsg(self_id, group_id, message, AutoRetract)
	}
}
