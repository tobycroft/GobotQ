package Group

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/service"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
)

type Interface struct {
	Db gorose.IOrm
}

func App_check_balance(self_id, group_id, user_id, message_id int64, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	auto_retract := false
	if groupfunction["sign_send_retract"].(int64) == 1 {
		auto_retract = true
	}
	go func(self_id, group_id, user_id, message_id int64, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
		if groupfunction["sign_send_retract"].(int64) == 1 {
			var ret api.Struct_Retract
			ret.MessageId = message_id
			ret.Self_id = self_id
			api.Retract_chan <- ret
		}
	}(self_id, group_id, user_id, message_id, groupmember, groupfunction)
	var gpm GroupBalanceModel.Interface
	gpm.Db = tuuz.Db()
	gbl := gpm.Api_find(group_id, user_id)
	at := service.Serv_at(user_id)
	str := at + "您当前拥有" + Calc.Any2String(gbl["balance"]) + "分"
	go api.Sendgroupmsg(self_id, group_id, str, auto_retract)
}

func App_check_rank(self_id, group_id, user_id, message_id int64, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	auto_retract := false
	if groupfunction["sign_send_retract"].(int64) == 1 {
		auto_retract = true
	}
	go func(self_id, group_id, user_id, message_id int64, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
		if groupfunction["sign_send_retract"].(int64) == 1 {
			var ret api.Struct_Retract
			ret.MessageId = message_id
			ret.Self_id = self_id
			api.Retract_chan <- ret
		}
	}(self_id, group_id, user_id, message_id, groupmember, groupfunction)
	var gpm GroupBalanceModel.Interface
	gpm.Db = tuuz.Db()
	gbl := gpm.Api_select(group_id, 10)
	str := ""
	for i1, i2 := range gbl {
		user := GroupMemberModel.Api_find(group_id, i2["user_id"].(int64))
		if len(user) > 0 {
			if len(Calc.Any2String(user["card"])) > 2 && Calc.Any2String(user["card"]) != "null" {
				str += "第" + Calc.Int2String(i1+1) + "名：" + user["card"].(string) + "，" + Calc.Any2String(i2["balance"]) + "分" + "\r\n"
			} else {
				str += "第" + Calc.Int2String(i1+1) + "名：" + user["nickname"].(string) + "，" + Calc.Any2String(i2["balance"]) + "分" + "\r\n"
			}
		}
	}
	go api.Sendgroupmsg(self_id, group_id, str, auto_retract)
}
