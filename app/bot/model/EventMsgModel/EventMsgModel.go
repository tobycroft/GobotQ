package EventMsgModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "event_msg"

func Api_insert(bot, uid, gid, operator, text, Type, subtype interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"bot":      bot,
		"uid":      uid,
		"gid":      gid,
		"operator": operator,
		"text":     text,
		"type":     Type,
		"subtype":  subtype,
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
