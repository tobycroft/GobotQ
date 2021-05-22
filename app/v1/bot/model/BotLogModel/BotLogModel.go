package BotLogModel

import (
	"github.com/gohouse/gorose/v2"
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

func Api_insert(self_id, log, data, Type interface{}) bool {
	db := tuuz.Db().Table(table)
	db.Data(map[string]interface{}{
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

func Api_select(self_id interface{}, page, limit int) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"self_id": self_id,
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

func Api_get(self_id, id interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
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
