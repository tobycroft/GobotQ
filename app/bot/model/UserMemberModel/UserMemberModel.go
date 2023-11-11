package UserMemberModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "user_member"

func Api_insert(qq, uname, password any) bool {
	db := tuuz.Db().Table(table)
	data := map[string]any{
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

func Api_find(qq any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_find_byQqandPassword(qq, password any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"qq":       qq,
		"password": password,
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

func Api_delete(qq any) bool {
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

func Api_update_uname(qq, uname any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"qq": qq,
	}
	db.Where(where)
	data := map[string]any{
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

func Api_update_password(qq, password any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"qq": qq,
	}
	db.Where(where)
	data := map[string]any{
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

func Api_update_all(qq, uname, password any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"qq": qq,
	}
	db.Where(where)
	data := map[string]any{
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
