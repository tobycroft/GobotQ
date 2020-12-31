package PrivateMsgModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "private_msg"

func Api_insert(bot, uid, text, req, seq, Type, subtype, file_id, file_md5, file_name, file_size interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"bot":       bot,
		"uid":       uid,
		"text":      text,
		"req":       req,
		"seq":       seq,
		"type":      Type,
		"subtype":   subtype,
		"file_id":   file_id,
		"file_md5":  file_md5,
		"file_name": file_name,
		"file_size": file_size,
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

func Api_select(bot interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot": bot,
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
