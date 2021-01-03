package Group

import (
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/app/bot/model/GroupSignModel"
)

func App_group_sign(gid, uid interface{}) {
	sign := GroupSignModel.Api_find(gid, uid)
	if len(sign) > 0 {

	} else {
		group_model := GroupBalanceModel.Api_find(gid, uid)
	}
}
