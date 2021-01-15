package GroupBalanceModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/config/app_conf"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_balance"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(gid, uid interface{}) bool {
	db := self.Db.Table(table)
	data := map[string]interface{}{
		"gid": gid,
		"uid": uid,
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

func Api_select(gid interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"gid": gid,
	}
	db.Where(where)
	db.Limit(app_conf.Db_default_load_limit)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_gt(gid, balance interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"gid": gid,
	}
	db.Where(where)
	db.Limit(app_conf.Db_default_load_limit)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_lt(gid, balance interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"gid": gid,
	}
	db.Where(where)
	db.Limit(app_conf.Db_default_load_limit)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_uid(uid interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"uid": uid,
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

func Api_find(gid, uid interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"gid": gid,
		"uid": uid,
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

func (self *Interface) Api_update(gid, uid, balance interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"gid": gid,
		"uid": uid,
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

func (self *Interface) Api_incr(gid, uid, balance_inc interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"gid": gid,
		"uid": uid,
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

func (self *Interface) Api_decr(gid, uid, balance_decr interface{}) bool {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"gid": gid,
		"uid": uid,
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
