package GroupAutoreplyModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_autoreply"

func Api_insert(gid, uid, key, value, percent interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"gid":     gid,
		"uid":     uid,
		"key":     key,
		"value":   value,
		"percent": percent,
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

func Api_select(gid, Type, percent interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"gid":  gid,
		"type": Type,
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

func Api_delete(gid, id interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"gid": gid,
		"id":  id,
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
