package UserBalanceModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "user_balance"

func Api_insert(qq, balance interface{}) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_insert(qq, balance)
}

func (self *Interface) Api_insert(qq, balance interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"qq":      qq,
		"balance": balance,
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

func Api_find(qq interface{}) gorose.Data {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_find(qq)
}

func (self *Interface) Api_find(qq interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"qq": qq,
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

type Interface struct {
	Db gorose.IOrm
}

func Api_dec_balance(qq interface{}, balance_dec float64) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_dec_balance(qq, balance_dec)
}

func (self *Interface) Api_dec_balance(qq interface{}, balance_dec float64) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"qq": qq,
	}
	db.Where(where)
	_, err := db.Decrement("balance", balance_dec)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
