package BotAdminModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "bot_admin"

func Api_insert(bot, qq, end_date any) bool {
	db := tuuz.Db().Table(table)
	data := map[string]any{
		"bot":      bot,
		"admin":    qq,
		"end_date": end_date,
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

func Api_delete(bot, qq any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"bot":   bot,
		"admin": qq,
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

func Api_inc_endTime(bot, qq any, endTime_incr int) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"bot":   bot,
		"admin": qq,
	}
	db.Where(where)
	_, err := db.Increment("endTime_incr", endTime_incr)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_select(qq any) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("id,cname,img,owner,admin,type,active,end_date")
	where := map[string]any{
		"admin": qq,
	}
	db.Where(where)
	db.Join("bot", "bot.bot=bot_admin.bot")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
