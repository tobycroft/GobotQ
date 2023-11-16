package TokenModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "token"

func Api_delete(uid interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"uid": uid,
	}
	db.Where(where)
	_, err := db.Delete()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_delete_byType(uid, Type interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"uid":  uid,
		"type": Type,
	}
	db.Where(where)
	_, err := db.Delete()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_insert(uid, token, Type, ip interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"uid":   uid,
		"token": token,
		"ip":    ip,
		"type":  Type,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find(uid, token interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"uid":   uid,
		"token": token,
	}
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_type(uid, token, Type interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"uid":   uid,
		"token": token,
		"type":  Type,
	}
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
func Api_delete_byId(qq, id any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"qq": qq,
		"id": id,
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
func Api_select(qq any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"qq": qq,
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
