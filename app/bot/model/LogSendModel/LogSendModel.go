package LogSendModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "log_send"

func Api_insert(self_id, Type, group_id, message any) bool {
	db := tuuz.Db().Table(Table)
	db.Data(map[string]any{
		"self_id":  self_id,
		"type":     Type,
		"group_id": group_id,
		"message":  message,
	})
	_, err := db.Insert()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
