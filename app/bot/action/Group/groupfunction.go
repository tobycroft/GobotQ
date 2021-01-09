package Group

import (
	"fmt"
	"main.go/app/bot/model/GroupFunctionDetailModel"
	"main.go/app/bot/model/GroupFunctionModel"
)

func App_group_function(bot *int, gid *int, uid *int) {
	fmt.Println(group_function_attach(*gid))
}

func group_function_attach(gid interface{}) map[string]interface{} {
	group_setting := GroupFunctionModel.Api_find(gid)
	if len(group_setting) < 1 {
		GroupFunctionModel.Api_insert(gid)
		return group_function_attach(gid)
	}
	function := GroupFunctionDetailModel.Api_select_kv()
	arr := make(map[string]interface{})
	for k, v := range group_setting {
		if function[k] != nil {
			function[k].(map[string]interface{})["value"] = v
			arr[k] = function[k]
		}
	}
	return arr
}
