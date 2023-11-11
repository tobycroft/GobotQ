package BotModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "bot"

type Interface struct {
	Db gorose.IOrm
}

func Api_insert(self_id, cname, Type, owner, secret, password, end_time, url any) bool {
	db := tuuz.Db()
	var self Interface
	self.Db = db
	return self.Api_insert(self_id, cname, Type, owner, secret, password, end_time, url)
}

func (self *Interface) Api_insert(self_id, cname, Type, owner, secret, password, end_time, url any) bool {
	db := self.Db.Table(table)
	data := map[string]any{
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
	db.Where("end_time>UNIX_TIMESTAMP(CURRENT_TIMESTAMP)")
	db.Where("active", "=", 1)
	db.Where("url IS NOT NULL")
	db.Where("url != ''")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find(self_id any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_find_byOwnerandBot(owner, self_id any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_find_byBot_WithoutPassword(self_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("self_id,cname,img,type,owner,end_time,active,date")
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

func Api_select_byOwner(owner any) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("self_id,cname,img,type,owner,end_time,active,date")
	where := map[string]any{
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

func Api_update_img(owner, self_id, img any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"owner":   owner,
		"self_id": self_id,
	}
	db.Where(where)
	data := map[string]any{
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

func Api_update_owner(self_id, owner any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"self_id": self_id,
	}
	db.Where(where)
	data := map[string]any{
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

func Api_update_secret(self_id, secret any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"self_id": self_id,
	}
	db.Where(where)
	data := map[string]any{
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

func Api_update_password(self_id, password any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"self_id": self_id,
	}
	db.Where(where)
	data := map[string]any{
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

func Api_update_cname(self_id, cname any) bool {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"self_id": self_id,
	}
	db.Where(where)
	data := map[string]any{
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

func Api_find_url(self_id any) gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("url")
	db.Where("self_id", self_id)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_update_url(self_id, url any) bool {
	db := tuuz.Db().Table(table)
	db.Fields("url")
	db.Where("self_id", self_id)
	db.Data(map[string]any{
		"url": url,
	})
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
