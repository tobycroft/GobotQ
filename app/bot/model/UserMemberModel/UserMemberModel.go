package UserMemberModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "user_member"

func Api_insert(qq, uname, password interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"qq":       qq,
		"uname":    uname,
		"password": password,
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

func Api_find(qq interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"qq": qq,
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

func Api_find_byQqandPassword(qq, password interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"qq": qq,
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

func Api_delete(qq interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
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

func Api_update_uname(qq, uname interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"qq": qq,
	}
	db.Where(where)
	data := map[string]interface{}{
		"uname": uname,
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

func Api_update_password(qq, password interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"qq": qq,
	}
	db.Where(where)
	data := map[string]interface{}{
		"password": password,
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

func Api_update_all(qq, uname, password interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"qq": qq,
	}
	db.Where(where)
	data := map[string]interface{}{
		"uname":    uname,
		"password": password,
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
