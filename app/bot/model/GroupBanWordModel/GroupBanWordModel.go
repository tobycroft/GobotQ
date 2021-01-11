package GroupBanWordModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_ban_word"

func Api_insert(bot, gid, uid, word, mode, is_kick, is_ban, is_retract, share interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"bot":        bot,
		"gid":        gid,
		"uid":        uid,
		"word":       word,
		"mode":       mode,
		"is_kick":    is_kick,
		"is_ban":     is_ban,
		"is_retract": is_retract,
		"share":      share,
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

func Api_select(gid interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"gid": gid,
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

func Api_delete(gid, word interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"gid":  gid,
		"word": word,
	}
	db.Where(where)
	_, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_update(gid, word, is_kick, is_ban, is_retract interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"gid":  gid,
		"word": word,
	}
	db.Where(where)
	data := map[string]interface{}{
		"is_kick":    is_kick,
		"is_ban":     is_ban,
		"is_retract": is_retract,
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
