package UserTokenModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "user_token"

func Api_insert(qq, token, ip any) bool {
	db := tuuz.Db().Table(table)
	data := map[string]any{
		"qq":    qq,
		"token": token,
		"ip":    ip,
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

func Api_find_byToken(qq, token any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"qq":    qq,
		"token": token,
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

func Api_delete_byToken(qq, token any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"qq":    qq,
		"token": token,
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

func Api_tuncate() {
	db := tuuz.Db().Table(table)
	db.Force()
	db.Delete()
}
