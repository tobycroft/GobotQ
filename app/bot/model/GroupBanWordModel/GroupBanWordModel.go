package GroupBanWordModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_ban_word"

func Api_insert(self_id, group_id, user_id, word, mode, is_kick, is_ban, is_retract, share interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"self_id":    self_id,
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

func Api_select_byKV(group_id interface{}, key string, value interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
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

func Api_select(group_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
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

func Api_find(group_id, word interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
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

func Api_delete(group_id, word interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
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

func Api_update(group_id, word, is_kick, is_ban, is_retract interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"word":     word,
	}
	db.Where(where)
	data := map[string]interface{}{
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
