package GroupBanModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_ban"

func Api_insert(group_id, user_id interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
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

func Api_find(group_id, user_id interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
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

func Api_count(group_id, user_id interface{}) int64 {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
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

func Api_delete(group_id interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
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

func Api_delete_userId(group_id, user_id interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
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
