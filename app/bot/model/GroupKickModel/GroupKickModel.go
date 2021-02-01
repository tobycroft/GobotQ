package GroupKickModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_kick"

func Api_insert(bot, gid, uid, last_msg interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"bot":      bot,
		"gid":      gid,
		"uid":      uid,
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
