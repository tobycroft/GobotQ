package AutoSendModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
	"time"
)

const table = "group_auto_send"

func Api_insert(gid, uid, ident, msg, Type, sep, count, next_time interface{}) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_insert(gid, uid, ident, msg, Type, sep, count, next_time)
}

func (self *Interface) Api_insert(gid, uid, ident, msg, Type, sep, count, next_time interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"gid":       gid,
		"uid":       uid,
		"ident":     ident,
		"msg":       msg,
		"type":      Type,
		"sep":       sep,
		"count":     count,
		"next_time": next_time,
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

func Api_update(gid, id, ident, msg, Type, sep, count, next_time interface{}) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_update(gid, id, ident, msg, Type, sep, count, next_time)
}

func (self *Interface) Api_update(gid, id, ident, msg, Type, sep, count, next_time interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"gid":       gid,
		"ident":     ident,
		"msg":       msg,
		"type":      Type,
		"sep":       sep,
		"count":     count,
		"next_time": next_time,
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
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select(gid interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where("gid", "=", gid)
	db.Order("id desc")
	ret, err := db.Get()
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

func Api_dec_count(id interface{}) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_dec_count(id)
}

func (self *Interface) Api_dec_count(id interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
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

func Api_update_next_time(id, next_time interface{}) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_update_next_time(id, next_time)
}

func (self *Interface) Api_update_next_time(id, next_time interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"id": id,
	}
	db.Where(where)
	db.Data(map[string]interface{}{
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

func Api_delete(gid, id interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"gid": gid,
		"id":  id,
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
