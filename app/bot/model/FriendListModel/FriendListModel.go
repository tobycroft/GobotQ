package FriendListModel

import (
	"encoding/json"
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "friend_list"

func Api_delete(self_id any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_delete_byUid(self_id, user_id any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_find(user_id any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_find_struct[T FriendList](user_id any) T {
	db := tuuz.Db().Table(table)
	if user_id != nil {
		db.Where("user_id", user_id)
	}
	ret := T{}
	err := db.Scan(&ret)
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return T{}
	} else {
		return ret
	}
}

func Api_select_struct[T FriendList](user_id any) []T {
	db := tuuz.Db().Table(table)
	if user_id != nil {
		db.Where("user_id", user_id)
	}
	ret := []T{}
	err := db.Scan(&ret)
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return []T{}
	} else {
		return ret
	}
}

func Api_select(self_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

type FriendList struct {
	Id       int64  `json:"id"`
	SelfId   int64  `gorose:"self_id" redis:"self_id"`
	UserId   int64  `gorose:"user_id" redis:"user_id"`
	Nickname string `gorose:"nickname" redis:"nickname"`
	Remark   string `gorose:"remark" redis:"remark"`
}

func (i *FriendList) MarshalBinary() ([]byte, error) {
	return json.Marshal(i)
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
