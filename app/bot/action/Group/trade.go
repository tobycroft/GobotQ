package Group

import (
	"main.go/app/bot/service"
	"main.go/config/app_default"
)

func App_trade_center(self_id, group_id, user_id, message_id int64, message string, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	switch message {

	case "我的", "列表", "背包":
		str := list_my_daoju(group_id, user_id)
		AutoMessage(self_id, group_id, user_id, str, groupfunction)
		break

	case "中心", "商城", "商店":
		str := list_daoju()
		AutoMessage(self_id, group_id, user_id, str, groupfunction)
		break

	case "帮助":
		AutoMessage(self_id, group_id, user_id, app_default.Default_trade, groupfunction)
		break

	case "购买", "兑换":
		AutoMessage(self_id, group_id, user_id, app_default.Daoju_goumai, groupfunction)
		break

	default:
		str, has := service.Serv_text_match(message, []string{"购买", "兑换"})
		if has {
			str, err := buy_daoju(group_id, user_id, str)
			if err != nil {
				AutoMessage(self_id, group_id, user_id, err.Error(), groupfunction)
			} else {
				AutoMessage(self_id, group_id, user_id, str, groupfunction)
			}
		}
		send_to, has := service.Serv_text_match(message, []string{"赠送", "赠与", "送给"})
		if has {
			send, err := send_daoju(group_id, user_id, send_to)
			if err != nil {
				AutoMessage(self_id, group_id, user_id, err.Error(), groupfunction)
			} else {
				AutoMessage(self_id, group_id, user_id, send, groupfunction)
			}
		}
		break
	}
}
