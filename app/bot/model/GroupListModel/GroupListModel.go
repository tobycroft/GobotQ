package GroupListModel

import (
	"github.com/tobycroft/gorose"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_list"

type GroupList struct {
	Self_id          interface{} `gorose:"self_id"`
	Group_id         int64       `gorose:"group_id"`
	Group_name       string      `gorose:"group_name"`
	Group_memo       string      `gorose:"group_memo"`
	Member_count     int64       `gorose:"member_count"`
	Max_member_count int64       `gorose:"max_member_count"`
}

func Api_insert(gl GroupList) bool {
	db := tuuz.Db().Table(table)
	db.Data(gl)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_insert_more(gls []GroupList) bool {
	db := tuuz.Db().Table(table)
	db.Data(gls)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
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

func Api_select_InGid(group_id []interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.WhereIn("group_id", group_id)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find(group_id interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
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

func Api_delete_byGid(group_id interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
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

func Api_delete_byBotandGid(self_id, group_id interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"self_id":  self_id,
		"group_id": group_id,
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
