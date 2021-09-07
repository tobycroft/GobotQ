package BotFriendAllowModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "bot_friend_allow"

func Api_insert(bot, uid interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"bot": bot,
		"uid": uid,
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

func Api_find(bot, uid interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot": bot,
		"uid": uid,
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

func Api_delete(bot, uid interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot": bot,
		"uid": uid,
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

func Api_select(bot interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
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
