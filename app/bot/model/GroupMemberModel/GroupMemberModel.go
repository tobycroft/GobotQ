package GroupMemberModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_member"

type GroupMember struct {
	SelfId       any    `gorose:"self_id" redis:"self_id"`
	Card         string `gorose:"card" redis:"card"`
	GroupId      any    `gorose:"group_id" redis:"group_id"`
	JoinTime     int64  `gorose:"join_time" redis:"join_time"`
	LastSentTime int64  `gorose:"last_sent_time" redis:"last_sent_time"`
	Level        int64  `gorose:"level" redis:"level"`
	Nickname     string `gorose:"nickname" redis:"nickname"`
	Role         string `gorose:"role" redis:"role"`
	Title        string `gorose:"title" redis:"title"`
	UserId       int64  `gorose:"user_id" redis:"user_id"`
}

func Api_insert(gm GroupMember) bool {
	db := tuuz.Db().Table(table)
	db.Data(gm)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_insert_more(gms []GroupMember) bool {
	db := tuuz.Db().Table(table)
	db.Data(gms)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_update(group_id, user_id any, gm GroupMember) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	db.Data(gm)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_select(self_id, group_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"self_id":  self_id,
		"group_id": group_id,
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

func Api_select_byGroupId(group_id any, order string, limit, page int) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("*", "FROM_UNIXTIME(last_sent_time) as last_date")
	where := map[string]any{
		"group_id": group_id,
		"role":     "member",
	}
	db.Where(where)
	db.OrderBy(order)
	db.Limit(limit)
	db.Page(page)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_count_byGroupIdAndRole(group_id, role any) int64 {
	db := tuuz.Db().Table(table)
	if group_id != nil {
		db.Where("group_id", group_id)
	}
	if role != nil {
		db.Where("role", role)
	}
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_select_groupBy_groupId(self_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"self_id": self_id,
	}
	db.GroupBy("group_id")
	db.Where(where)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_byUid(user_id any, role []any) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where("user_id", user_id)
	db.WhereIn("role", role)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_delete_byGid(self_id, group_id any) bool {
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

func Api_delete_byUid(self_id, group_id, user_id any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"user_id":  user_id,
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

func Api_find(group_id, user_id any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
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

func Api_find_byUid(user_id any) gorose.Data {
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

func Api_update_type(group_id, user_id, role any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	data := map[string]any{
		"role": role,
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

func Api_find_owner(self_id, group_id any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"self_id":  self_id,
		"group_id": group_id,
		"role":     "owner",
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

func Api_select_admin(self_id, group_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"self_id":  self_id,
		"group_id": group_id,
		"role":     "admin",
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

func Api_find_struct[T GroupMember](self_id, user_id, group_id any) T {
	db := tuuz.Db().Table(table)
	if self_id != nil {
		db.Where("self_id", self_id)
	}
	if user_id != nil {
		db.Where("user_id", user_id)
	}
	if group_id != nil {
		db.Where("group_id", group_id)
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

func Api_select_struct[T GroupMember](self_id, user_id, group_id any) []T {
	db := tuuz.Db().Table(table)
	if self_id != nil {
		db.Where("self_id", self_id)
	}
	if user_id != nil {
		db.Where("user_id", user_id)
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
