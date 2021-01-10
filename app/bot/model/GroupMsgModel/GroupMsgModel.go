package GroupMsgModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_msg"

func Api_insert(bot, uid, gid, text, req, random, file_id, file_md5, file_name, file_size, send, recv interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"bot":       bot,
		"uid":       uid,
		"gid":       gid,
		"text":      text,
		"req":       req,
		"random":    random,
		"file_id":   file_id,
		"file_md5":  file_md5,
		"file_name": file_name,
		"file_size": file_size,
		"send":      send,
		"recv":      recv,
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

func Api_find(bot, gid, send interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot":  bot,
		"gid":  gid,
		"send": send,
	}
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
