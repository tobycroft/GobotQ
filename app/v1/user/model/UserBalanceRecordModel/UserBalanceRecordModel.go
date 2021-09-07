package UserBalanceRecordModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "user_balance_record"

type Interface struct {
	Db gorose.IOrm
}

func Api_insert(qq, before_balance, amount, after_balance, remark interface{}) bool {
	var self Interface
	self.Db = tuuz.Db()
	return self.Api_insert(qq, before_balance, amount, after_balance, remark)
}

func (self *Interface) Api_insert(qq, before_balance, amount, after_balance, remark interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"qq":             qq,
		"before_balance": before_balance,
		"amount":         amount,
		"after_balance":  after_balance,
		"remark":         remark,
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

func Api_select(qq interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"qq": qq,
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
	db.LockForUpdate()
	db.Order("id desc")
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
