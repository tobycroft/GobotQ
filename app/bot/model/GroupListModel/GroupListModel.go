package GroupListModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_list"

type GroupList struct {
	Id             int64  `json:"id,omitempty"`
	SelfId         int64  `gorose:"self_id" redis:"self_id" json:"self_id,omitempty"`
	GroupId        int64  `gorose:"group_id" redis:"group_id" json:"group_id,omitempty"`
	GroupName      string `gorose:"group_name" redis:"group_name" json:"group_name,omitempty"`
	GroupMemo      string `gorose:"group_memo" redis:"group_memo" json:"group_memo,omitempty"`
	MemberCount    int64  `gorose:"member_count" redis:"member_count" json:"member_count,omitempty"`
	MaxMemberCount int64  `gorose:"max_member_count" redis:"max_member_count" json:"max_member_count,omitempty"`
	Admins         string `gorose:"admins" redis:"admins" json:"admins,omitempty"`
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

func Api_select_InGid(group_id []any) []gorose.Data {
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

func Api_find(group_id any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_delete_byGid(group_id any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_delete_byBotandGid(self_id, group_id any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_find_struct(self_id, group_id any) GroupList {
	db := tuuz.Db().Table(table)
	if self_id != nil {
		db.Where("self_id", self_id)
	}
	if group_id != nil {
		db.Where("group_id", group_id)
	}
	ret := GroupList{}
	err := db.Scan(&ret)
	if err != nil {
		Log.DBrrsql(err, db, tuuz.FUNCTION_ALL())
		return GroupList{}
	} else {
		return ret
	}
}

func Api_select_struct[T GroupList](self_id, group_id any) []T {
	db := tuuz.Db().Table(table)
	if self_id != nil {
		db.Where("self_id", self_id)
	}
	if group_id != nil {
		db.Where("group_id", group_id)
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
