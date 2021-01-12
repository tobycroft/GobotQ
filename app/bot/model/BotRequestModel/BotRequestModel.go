package BotRequestModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "bot_request"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(uid, bot, password, owner, secret interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"uid":      uid,
		"bot":      bot,
		"password": password,
		"owner":    owner,
		"secret":   secret,
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

func Api_find(bot, password interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot":      bot,
		"password": password,
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

func (self *Interface) Api_delete(bot interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"bot": bot,
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
