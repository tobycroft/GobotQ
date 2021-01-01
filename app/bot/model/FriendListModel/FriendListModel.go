package FriendListModel

import (
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
	bot      interface{}
	uid      interface{}
	nickname interface{}
	remark   interface{}
	email    interface{}
}

type FriendLists struct {
	Fl []FriendList
}

func Api_insert_more(fl FriendLists) bool {
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
