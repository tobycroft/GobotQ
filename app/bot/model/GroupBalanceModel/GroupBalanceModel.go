package GroupBalanceModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_balance"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(group_id, user_id interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
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

func Api_select(group_id interface{}, limit int) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
	}
	db.Where(where)
	db.Order("balance desc")
	db.Limit(limit)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_gt(group_id, balance interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
	}
	db.Where(where)
	db.Where("balance", ">", balance)
	db.Order("balance asc")
	db.Limit(9)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_lt(group_id, balance interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
	}
	db.Where(where)
	db.Where("balance", "<", balance)
	db.Order("balance desc")
	db.Limit(9)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_uid(user_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"user_id": user_id,
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

func Api_find(group_id, user_id interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
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

func (self *Interface) Api_update(group_id, user_id, balance interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	data := map[string]interface{}{
		"balance": balance,
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

func (self *Interface) Api_incr(group_id, user_id, balance_inc interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	_, err := db.Increment("balance", balance_inc)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_decr(group_id, user_id, balance_decr interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	_, err := db.Decrement("balance", balance_decr)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_value_balance(group_id, user_id interface{}) interface{} {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	ret, err := db.Value("balance")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
