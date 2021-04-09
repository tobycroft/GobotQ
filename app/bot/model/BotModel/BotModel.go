package BotModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "self_id"

type Interface struct {
	Db gorose.IOrm
}

func Api_insert(self_id, cname, Type, owner, secret, password, end_time, url interface{}) bool {
	db := tuuz.Db()
	var self Interface
	self.Db = db
	return self.Api_insert(self_id, cname, Type, owner, secret, password, end_time, url)
}

func (self *Interface) Api_insert(self_id, cname, Type, owner, secret, password, end_time, url interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"self_id":  self_id,
		"cname":    cname,
		"Type":     Type,
		"owner":    owner,
		"secret":   secret,
		"password": password,
		"end_time": end_time,
		"url":      url,
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

func Api_select() []gorose.Data {
	db := tuuz.Db().Table(table)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find(self_id interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"self_id": self_id,
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

func Api_find_byOwnerandBot(owner, self_id interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"owner":   owner,
		"self_id": self_id,
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

func Api_find_byBot_WithoutPassword(self_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("self_id,cname,img,type,owner,end_time,active,date")
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

func Api_select_byOwner(owner interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("self_id,cname,img,type,owner,end_time,active,date")
	where := map[string]interface{}{
		"owner": owner,
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

func Api_update_img(owner, self_id, img interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"owner":   owner,
		"self_id": self_id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"img": img,
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

func Api_update_owner(self_id, owner interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"self_id": self_id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"owner": owner,
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

func Api_update_secret(self_id, secret interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"self_id": self_id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"secret": secret,
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

func Api_update_password(self_id, password interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"self_id": self_id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"password": password,
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

func Api_update_cname(self_id, cname interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"self_id": self_id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"cname": cname,
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
