package Group

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupFunctionDetailModel"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/tuuz/Calc"
)

func App_group_function(bot *int, gid *int, uid *int) {
	settings := group_function_attach(*gid)
	str := "您的群设定为：\r\n"
	for _, v := range settings {
		str += Calc.Any2String(v["name"]) + ":"
		value := Calc.Any2String(v["value"])
		switch v["type"] {
		case "bool":
			if value == "1" {
				str += "是"
			} else {
				str += "否"
			}
			break

		case "int":
			str += value
			break

		case "string":
			str += value
			break

		default:
			str += value
			break
		}

		str += "\r\n"
	}
	api.Sendgroupmsg(*bot, *gid, str)
}

func group_function_attach(gid interface{}) map[string]map[string]interface{} {
	group_setting := GroupFunctionModel.Api_find(gid)
	if len(group_setting) < 1 {
		GroupFunctionModel.Api_insert(gid)
		return group_function_attach(gid)
	}
	function := GroupFunctionDetailModel.Api_select_kv()
	arr := make(map[string]map[string]interface{})
	for k, v := range group_setting {
		if function[k] != nil {
			function[k]["value"] = v
			arr[k] = function[k]
		}
	}
	return arr
}
