package GroupBanWordModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_ban_word"

func Api_insert(group_id, user_id, word, mode, is_kick, is_ban, is_retract, share any) bool {
	db := tuuz.Db().Table(table)
	data := map[string]any{
		"group_id":   group_id,
		"user_id":    user_id,
		"word":       word,
		"mode":       mode,
		"is_kick":    is_kick,
		"is_ban":     is_ban,
		"is_retract": is_retract,
		"share":      share,
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

func Api_select_byKV(group_id any, key string, value any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		key:        value,
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

func Api_find(group_id, word any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"word":     word,
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

func Api_delete(group_id, word any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"word":     word,
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
func Api_delete_byId(group_id, id any) bool {
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

func Api_update(group_id, word, is_kick, is_ban, is_retract any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"word":     word,
	}
	db.Where(where)
	data := map[string]any{
		"is_kick":    is_kick,
		"is_ban":     is_ban,
		"is_retract": is_retract,
	}
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
