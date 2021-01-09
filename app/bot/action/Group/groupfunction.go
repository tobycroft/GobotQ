package Group

import (
	"fmt"
	"github.com/gohouse/gorose/v2"
	"main.go/app/bot/model/GroupFunctionDetailModel"
	"main.go/app/bot/model/GroupFunctionModel"
)

func App_group_function(bot *int, gid *int, uid *int) {
	fmt.Println(group_function_attach(*gid))
}

func group_function_attach(gid interface{}) gorose.Data {
	group_setting := GroupFunctionModel.Api_find(gid)
	function := GroupFunctionDetailModel.Api_select_kv()
	for k, v := range group_setting {
		function[k].(gorose.Data)["value"] = v
		group_setting[k] = function[k]
	}
	return group_setting
}
