package GroupDaojuModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_daoju"

type Interface struct {
	Db gorose.IOrm
}

func Api_select(group_id, user_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
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

func Api_select_have(group_id, user_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
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

func Api_find(group_id, user_id, dj_id interface{}) gorose.Data {
	var self Interface
	self.Db = tuuz.Db()
	return Api_find(group_id, user_id, dj_id)
}

func (self *Interface) Api_find(group_id, user_id, dj_id interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
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

func Api_value(group_id, user_id, dj_id interface{}) interface{} {
	var self Interface
	self.Db = tuuz.Db()
	return Api_value(group_id, user_id, dj_id)
}

func (self *Interface) Api_value(group_id, user_id, dj_id interface{}) interface{} {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
	}
	db.Where(where)
	data, err := db.Value("num")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return data
	}
}

func (self *Interface) Api_decr(group_id, user_id, dj_id interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
	}
	db.Where(where)
	_, err := db.Decrement("num", 1)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_incr(group_id, user_id, dj_id interface{}, num int64) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
		"dj_id":    dj_id,
	}
	db.Where(where)
	_, err := db.Increment("num", num)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_insert(group_id, user_id, dj_id, num interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
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

func (self *Interface) Api_delete(group_id, user_id interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
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
