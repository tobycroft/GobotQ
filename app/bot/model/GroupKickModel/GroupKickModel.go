package GroupKickModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_kick"

func Api_insert(self_id, group_id, user_id, last_msg any) bool {
	db := tuuz.Db().Table(table)
	data := map[string]any{
		"self_id":  self_id,
		"group_id": group_id,
		"user_id":  user_id,
		"last_msg": last_msg,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_select(group_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
	}
	db.Where(where)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
