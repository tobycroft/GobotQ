package GroupBlackListModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_black_list"

func Api_insert(group_id, user_id, operator any) bool {
	db := tuuz.Db().Table(table)
	data := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
		"operator": operator,
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

func Api_find(group_id, user_id any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_count(group_id, user_id any) int64 {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_delete(group_id, user_id any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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
