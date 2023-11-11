package UserBalanceModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "user_balance"

func Api_insert(qq, balance any) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_insert(qq, balance)
}

func (self *Interface) Api_insert(qq, balance any) bool {
	db := self.Db.Table(table)
	data := map[string]any{
		"qq":      qq,
		"balance": balance,
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

func Api_find_balance(qq any) any {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_find_balance(qq)
}

func (self *Interface) Api_find_balance(qq any) float64 {
	db := self.Db.Table(table)
	where := map[string]any{
		"qq": qq,
	}
	db.Where(where)
	db.LockForUpdate()
	ret, err := db.Value("balance")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret.(float64)
	}
}

func Api_find(qq any) gorose.Data {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_find(qq)
}

func (self *Interface) Api_find(qq any) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]any{
		"qq": qq,
	}
	db.Where(where)
	db.LockForUpdate()
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

func Api_dec_balance(qq any, balance_dec float64) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_dec_balance(qq, balance_dec)
}

func (self *Interface) Api_dec_balance(qq any, balance_dec float64) bool {
	db := self.Db.Table(table)
	where := map[string]any{
		"qq": qq,
	}
	db.Where(where)
	db.LockForUpdate()
	_, err := db.Decrement("balance", balance_dec)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_inc_balance(qq any, balance_inc float64) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_inc_balance(qq, balance_inc)
}

func (self *Interface) Api_inc_balance(qq any, balance_inc float64) bool {
	db := self.Db.Table(table)
	where := map[string]any{
		"qq": qq,
	}
	db.Where(where)
	db.LockForUpdate()
	_, err := db.Increment("balance", balance_inc)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
