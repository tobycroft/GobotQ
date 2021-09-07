package GroupFunctionModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_function"

func Api_insert(group_id interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"group_id": group_id,
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

func Api_find(group_id interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
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

func Api_update(group_id, key, value interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
	}
	db.Where(where)
	data := map[interface{}]interface{}{
		key: value,
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
