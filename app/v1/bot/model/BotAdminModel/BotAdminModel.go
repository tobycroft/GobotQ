package BotAdminModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "bot_admin"

func Api_insert(bot, qq, end_time interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"bot":      bot,
		"admin":    qq,
		"end_time": end_time,
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

func Api_delete(bot, qq interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
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

func Api_inc_endTime(bot, qq interface{}, endTime_incr int) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
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

func Api_select(bot, qq interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot":   bot,
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
