package PrivateMsgModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "private_msg"

func Api_insert(self_id, user_id, message_id, message, raw_message, sub_type, time interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"self_id":     self_id,
		"user_id":     user_id,
		"message_id":  message_id,
		"message":     message,
		"raw_message": raw_message,
		"sub_type":    sub_type,
		"time":        time,
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

func Api_select(self_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"self_id": self_id,
	}
	db.Where(where)
	db.Limit(30)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
