package PrivateAutoReplyModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "private_auto_reply"

func Api_insert(self_id, qq, mode, key, value any) bool {
	db := tuuz.Db().Table(table)
	data := map[string]any{
		"self_id": self_id,
		"qq":      qq,
		"mode":    mode,
		"key":     key,
		"value":   value,
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

func Api_find_byKey(key any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"key": key,
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

func Api_select_byKey(key any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"key": key,
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

func Api_select_semi(self_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"self_id": self_id,
		"mode":    "semi",
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

func Api_delete(self_id, id any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"self_id": self_id,
		"id":      id,
	}
	db.Where(where)
	_, err := db.Delete()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_delete_byQq(qq any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"qq": qq,
	}
	db.Where(where)
	_, err := db.Delete()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
