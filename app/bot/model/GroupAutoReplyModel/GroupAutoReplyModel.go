package GroupAutoReplyModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_auto_reply"

func Api_insert(Type, group_id, user_id, key, value, percent any) bool {
	db := tuuz.Db().Table(table)
	data := map[string]any{
		"type":     Type,
		"group_id": group_id,
		"user_id":  user_id,
		"key":      key,
		"value":    value,
		"percent":  percent,
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

func Api_select(group_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
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

func Api_select_semi(group_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"type":     "semi",
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

func Api_select_semi_byPercent(group_id, percent any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"type":     "semi",
	}
	db.Where(where)
	db.Where("percent", ">", percent)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_full(group_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"type":     "full",
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

func Api_find(group_id, key, percent any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"key":      key,
		"percent":  percent,
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

func Api_delete(group_id, id any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"id":       id,
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
