package GroupDaojuModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_daoju"

type Interface struct {
	Db gorose.IOrm
}

func Api_select(group_id, user_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	data, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return data
	}
}

func Api_select_have(group_id, user_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	db.Where("num", ">", 0)
	data, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return data
	}
}

func Api_find(group_id, user_id, dj_id any) gorose.Data {
	var self Interface
	self.Db = tuuz.Db()
	return Api_find(group_id, user_id, dj_id)
}

func (self *Interface) Api_find(group_id, user_id, dj_id any) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
	}
	db.Where(where)
	data, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return data
	}
}

func Api_value(group_id, user_id, dj_id any) any {
	var self Interface
	self.Db = tuuz.Db()
	return Api_value(group_id, user_id, dj_id)
}

func (self *Interface) Api_value_num(group_id, user_id, dj_id any) int64 {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
	}
	db.Where(where)
	db.LockForUpdate()
	data, err := db.Value("num")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		if data == nil {
			return 0
		} else {
			return data.(int64)
		}
	}
}

func (self *Interface) Api_decr(group_id, user_id, dj_id any) bool {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
	}
	db.Where(where)
	db.LockForUpdate()
	_, err := db.Decrement("num", 1)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_incr(group_id, user_id, dj_id any, num int64) bool {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
	}
	db.Where(where)
	db.LockForUpdate()
	_, err := db.Increment("num", num)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_insert(group_id, user_id, dj_id, num any) bool {
	db := self.Db.Table(table)
	data := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
		"num":      num,
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

func (self *Interface) Api_delete(group_id, user_id any) bool {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
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

func (self *Interface) Api_find_in_djId(group_id, user_id any, dj_id []any) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	db.WhereIn("dj_id", dj_id)
	db.Where("num", ">", 0)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
