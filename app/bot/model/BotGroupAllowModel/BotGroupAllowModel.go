package BotGroupAllowModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "bot_group_allow"

func Api_insert(bot, gid any) bool {
	db := tuuz.Db().Table(table)
	data := map[string]any{
		"bot": bot,
		"gid": gid,
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

func Api_find(bot, gid any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"bot": bot,
		"gid": gid,
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

func Api_delete(bot, gid any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"bot": bot,
		"gid": gid,
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

func Api_select(bot any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"bot": bot,
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
