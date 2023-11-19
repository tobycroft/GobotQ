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

func Api_insert(self_id, cname, Type, owner, secret, password, end_date, url any) bool {
	db := tuuz.Db()
	var self Interface
	self.Db = db
	return self.Api_insert(self_id, cname, Type, owner, secret, password, end_date, url)
}

func (self *Interface) Api_insert(self_id, cname, Type, owner, secret, password, end_date, url any) bool {
	db := self.Db.Table(table)
	data := map[string]any{
		"self_id":  self_id,
		"cname":    cname,
		"Type":     Type,
		"owner":    owner,
		"secret":   secret,
		"password": password,
		"end_date": end_date,
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
	db.Where("end_date>UNIX_TIMESTAMP(CURRENT_TIMESTAMP)")
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
	db.Fields("self_id,cname,img,type,owner,end_date,active,date")
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
	db.Fields("self_id,cname,img,type,owner,end_date,active,date")
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

func Api_update_allowIp(self_id, allow_ip any) bool {
	db := tuuz.Db().Table(table)
	db.Where("self_id", self_id)
	db.Data(map[string]any{
		"allow_ip": allow_ip,
	})
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_update_active(self_id, active any) bool {
	db := tuuz.Db().Table(table)
	db.Where("self_id", self_id)
	db.Data(map[string]any{
		"active": active,
	})
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_update_manual(self_id, data any) bool {
	db := tuuz.Db().Table(table)
	db.Where("self_id", self_id)
	db.Data(data)
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_select_public(Type any) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("self_id,cname,img,type,owner,end_date,active,date")
	db.Where("type", Type)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
