package DaojuModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "daoju"

func Api_select() []gorose.Data {
	db := tuuz.Db().Table(table)
	data, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return data
	}
}

func Api_select_canShow() []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"can_show": 1,
	}
	db.Where(where)
	data, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return data
	}
}

func Api_find(id any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"id": id,
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

func Api_find_canUse(id any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"id":      id,
		"can_use": true,
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

func Api_find_byCname(cname any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"cname": cname,
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

func Api_find_byName(name any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"name": name,
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
