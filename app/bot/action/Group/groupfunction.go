package Group

import (
	"main.go/app/bot/model/GroupFunctionDetailModel"
	"main.go/app/bot/model/GroupFunctionModel"
)

func App_group_function(bot, gid, uid interface{}) {
	group_setting := GroupFunctionModel.Api_find(gid)
	function := GroupFunctionDetailModel.Api_select_kv()
	for k, v := range group_setting {
		function[k].(map[string]interface{})["value"] = v
		group_setting[k] = function[k]
	}
}
