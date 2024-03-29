package GroupBalanceModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_balance"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(group_id, user_id any) bool {
	db := self.Db.Table(table)
	data := map[string]any{
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

func (self *Interface) Api_select(group_id any, limit int) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_select_gt(group_id, balance any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_select_lt(group_id, balance any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func Api_select_uid(user_id any) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
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

func (self *Interface) Api_find(group_id, user_id any) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
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

func (self *Interface) Api_count_gt_balance(group_id, balance any) int64 {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
	}
	db.Where("balance", ">", balance)
	db.Where(where)
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func (self *Interface) Api_update(group_id, user_id, balance any) bool {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	data := map[string]any{
		"balance": balance,
	}
	db.Data(data)
	db.LockForUpdate()
	_, err := db.Update()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_incr(group_id, user_id, balance_inc any) bool {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
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

func (self *Interface) Api_decr(group_id, user_id, balance_decr any) bool {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	db.LockForUpdate()
	_, err := db.Decrement("balance", balance_decr)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_value_balance(group_id, user_id any) any {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	db.LockForUpdate()
	ret, err := db.Value("balance")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
