package BotSettingModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "bot_setting"

func Api_insert(self_id, add_friend, add_group any) bool {
	db := tuuz.Db().Table(Table)
	db.Data(map[string]any{
		"self_id":    self_id,
		"add_friend": add_friend,
		"add_group":  add_group,
	})
	_, err := db.Insert()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_update(self_id, data any) bool {
	db := tuuz.Db().Table(Table)
	db.Where("self_id", self_id)
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find(self_id any) gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Where("self_id", self_id)
	ret, err := db.Find()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
