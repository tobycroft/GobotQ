package GroupMsgModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_msg"

func Api_insert(self_id, user_id, group_id, message, raw_message, message_id, sub_type interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"self_id":     self_id,
		"user_id":     user_id,
		"group_id":    group_id,
		"message":     message,
		"raw_message": raw_message,
		"message_id":  message_id,
		"sub_type":    sub_type,
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

func Api_select(group_id, user_id interface{}, limit int) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	db.Order("id desc")
	db.Limit(limit)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
