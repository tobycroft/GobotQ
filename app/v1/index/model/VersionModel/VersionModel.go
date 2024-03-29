package VersionModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "version"

func Api_find(platform any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"platform": platform,
	}
	db.Where(where)
	db.Order("id desc")
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select(platform any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"platform": platform,
	}
	db.Where(where)
	db.Order("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
