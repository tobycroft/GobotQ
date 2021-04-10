package FriendListModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "friend_list"

func Api_delete(self_id interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"self_id": self_id,
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

func Api_delete_byUid(self_id, user_id interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"self_id": self_id,
		"user_id": user_id,
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

func Api_find(user_id interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"user_id": user_id,
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

func Api_select(self_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"self_id": self_id,
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

func Api_insert(self_id, user_id, nickname interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"self_id":  self_id,
		"user_id":  user_id,
		"nickname": nickname,
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
	SelfId   interface{} `gorose:"self_id"`
	UserId   interface{} `gorose:"user_id"`
	Nickname string      `gorose:"nickname"`
	Remark   string      `gorose:"remark"`
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
