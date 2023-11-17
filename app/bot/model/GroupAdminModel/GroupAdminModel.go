package GroupAdminModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const Table = "group_admin"

type GroupAdmins struct {
	SelfId  any   `gorose:"self_id" redis:"self_id"`
	GroupId int64 `gorose:"group_id" redis:"group_id"`
	UserId  int64 `gorose:"user_id" redis:"user_id"`
}

func Api_insert_more(data []GroupAdmins) bool {
	db := tuuz.Db().Table(Table)
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
func Api_insert(self_id, group_id, user_id any) bool {
	db := tuuz.Db().Table(Table)
	db.Data(map[string]any{
		"self_id":  self_id,
		"group_id": group_id,
		"user_id":  user_id,
	})
	_, err := db.Insert()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_delete_bySelfIdAndGroupId(self_id, group_id any) bool {
	db := tuuz.Db().Table(Table)
	if group_id != nil {
		db.Where("group_id", group_id)
	}
	if self_id != nil {
		db.Where("self_id", self_id)
	}
	_, err := db.Delete()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find_byGroupIdAndUserId(group_id, user_id any) gorose.Data {
	db := tuuz.Db().Table(Table)
	db.Where("group_id", group_id)
	db.Where("user_id", user_id)
	ret, err := db.Find()
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
