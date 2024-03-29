package AutoSendModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"time"
)

const table = "group_auto_send"

func Api_insert(group_id, user_id, ident, msg, Type, sep, count, next_time, retract any) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_insert(group_id, user_id, ident, msg, Type, sep, count, next_time, retract)
}

func (self *Interface) Api_insert(group_id, user_id, ident, msg, Type, sep, count, next_time, retract any) bool {
	db := self.Db.Table(table)
	data := map[string]any{
		"group_id":  group_id,
		"user_id":   user_id,
		"ident":     ident,
		"msg":       msg,
		"type":      Type,
		"sep":       sep,
		"count":     count,
		"next_time": next_time,
		"retract":   retract,
	}
	db.Data(data)
	db.LockForUpdate()
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_update(group_id, user_id, id, ident, msg, Type, sep, count, next_time, retract any) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_update(group_id, user_id, id, ident, msg, Type, sep, count, next_time, retract)
}

func (self *Interface) Api_update(group_id, user_id, id, ident, msg, Type, sep, count, next_time, retract any) bool {
	db := self.Db.Table(table)
	where := map[string]any{
		"id": id,
	}
	db.Where(where)
	data := map[string]any{
		"group_id":  group_id,
		"user_id":   user_id,
		"ident":     ident,
		"msg":       msg,
		"type":      Type,
		"sep":       sep,
		"count":     count,
		"next_time": next_time,
		"retract":   retract,
	}
	db.Data(data)
	db.LockForUpdate()
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_select_next_time_up() []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where("next_time", "<", time.Now().Unix())
	db.Where("count", ">", 0)
	db.Where("active", "=", 1)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select(group_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where("group_id", "=", group_id)
	db.Order("id desc")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find(group_id, id any) gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where("group_id", "=", group_id)
	db.Where("id", "=", id)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

type Interface struct {
	Db gorose.IOrm
}

func Api_dec_count(id any) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_dec_count(id)
}

func (self *Interface) Api_dec_count(id any) bool {
	db := self.Db.Table(table)
	where := map[string]any{
		"id": id,
	}
	db.Where(where)
	db.LockForUpdate()
	_, err := db.Decrement("count", 1)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_update_next_time(group_id, id, next_time any) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_update_next_time(group_id, id, next_time)
}

func (self *Interface) Api_update_next_time(group_id, id, next_time any) bool {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"id":       id,
	}
	db.Where(where)
	db.Data(map[string]any{
		"next_time": next_time,
	})
	db.LockForUpdate()
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_delete(group_id, id any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"id":       id,
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

func Api_update_active(group_id, id, active any) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_update_active(group_id, id, active)
}

func (self *Interface) Api_update_active(group_id, id, active any) bool {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"id":       id,
	}
	db.Where(where)
	db.Data(map[string]any{
		"active": active,
	})
	db.LockForUpdate()
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
