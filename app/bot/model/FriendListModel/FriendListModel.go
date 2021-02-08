package FriendListModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "friend_list"

func Api_delete(bot interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot": bot,
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

func Api_delete_byUid(bot, uid interface{}) bool {
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

func Api_find(uid interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
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

func Api_insert(bot, uid, nickname, remark, email interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"bot":      bot,
		"uid":      uid,
		"nickname": nickname,
		"remark":   remark,
		"email":    email,
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

type FriendList struct {
	Bot      interface{} `gorose:"bot"`
	Uid      interface{} `gorose:"uid"`
	Nickname string      `gorose:"nickname"`
	Remark   string      `gorose:"remark"`
	Email    string      `gorose:"email"`
}

func Api_insert_more(fl []FriendList) bool {
	db := tuuz.Db().Table(table)
	db.Data(fl)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
