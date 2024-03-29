package GroupFunctionDetailModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_function_detail"

func Api_find_byK(key any) any {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"key": key,
	}
	db.Where(where)
	ret, err := db.Value("name")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_type_byName(name any) any {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"name": name,
	}
	db.Where(where)
	ret, err := db.Value("type")
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

func Api_select() []gorose.Data {
	db := tuuz.Db().Table(table)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_kv() map[string]map[string]any {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"show": true,
	}
	db.Where(where)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		datas := map[string]map[string]any{}
		for _, data := range ret {
			datas[data["key"].(string)] = map[string]any{
				"name": data["name"],
				"type": data["type"],
				"show": data["show"],
			}
		}
		return datas
	}
}
