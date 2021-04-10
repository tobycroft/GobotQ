package Group

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBanModel"
	"main.go/tuuz/Calc"
	"math"
)

func App_ban_user(self_id, group_id, user_id interface{}, auto_retract bool, groupfunction map[string]interface{}) {
	time := GroupBanModel.Api_count(group_id, user_id)
	GroupBanModel.Api_insert(group_id, user_id)
	api.Sendgroupmsg(self_id, group_id, "这是你第"+Calc.Any2String(time+1)+"次接受惩罚，积分扣除", auto_retract)
	api.SetGroupBan(self_id, group_id, user_id, float64(groupfunction["ban_time"].(int64))*math.Pow10(int(time)))
}

func Api_kick_user(bot, gid, uid int, req int, random int, groupmember map[string]interface{}, groupfunction map[string]interface{}) {

}

func Api_retract_send(bot, gid, uid int, req int, random int, groupmember map[string]interface{}, groupfunction map[string]interface{}) {

}
