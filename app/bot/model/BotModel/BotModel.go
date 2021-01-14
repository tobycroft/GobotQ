package BotModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "bot"

type Interface struct {
	Db gorose.IOrm
}

func Api_insert(bot, cname, Type, owner, secret, password, end_time interface{}) bool {
	db := tuuz.Db()
	var self Interface
	self.Db = db
	return self.Api_insert(bot, cname, Type, owner, secret, password, end_time)
}

func (self *Interface) Api_insert(bot, cname, Type, owner, secret, password, end_time interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"bot":      bot,
		"cname":    cname,
		"Type":     Type,
		"owner":    owner,
		"secret":   secret,
		"password": password,
		"end_time": end_time,
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

func Api_find(bot interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot": bot,
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

func Api_find_byOwnerandBot(owner, bot interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"owner": owner,
		"bot":   bot,
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

func Api_select_byOwner(owner interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	db.Fields("bot,cname,img,type,owner,end_time,active,date")
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

func Api_update_img(owner, bot, img interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"owner": owner,
		"bot":   bot,
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

func Api_update_owner(bot, owner interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot": bot,
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

func Api_update_secret(bot, secret interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot": bot,
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

func Api_update_password(bot, password interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot": bot,
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

func Api_update_cname(bot, cname interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot": bot,
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
