package Group

import (
	"github.com/gohouse/gorose/v2"
	"main.go/app/bot/model/GroupFunctionDetailModel"
	"main.go/app/bot/model/GroupFunctionModel"
)

func App_group_function(bot, gid, uid interface{}) {
	group_setting := GroupFunctionModel.Api_find(gid)
	function := GroupFunctionDetailModel.Api_select_kv()
	for k, v := range group_setting {
		function[k].(gorose.Data)["value"] = v
		group_setting[k] = function[k]
	}
}
