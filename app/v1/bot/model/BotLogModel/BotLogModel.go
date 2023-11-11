package BotLogModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "bot_log"

/*
table:bot_log

collection:

self_id
log
data
type
date
*/

func Api_insert(self_id, log, data, Type any) bool {
	db := tuuz.Db().Table(table)
	db.Data(map[string]any{
		"log":     log,
		"self_id": self_id,
		"data":    data,
		"type":    Type,
	})
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}

}

func Api_select(self_id any, page, limit int) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"self_id": self_id,
	}
	db.Where(where)
	db.Limit(limit)
	db.Page(page)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_get(self_id, id any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"self_id": self_id,
		"id":      id,
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
