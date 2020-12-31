package FriendListModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "friend_list"

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
