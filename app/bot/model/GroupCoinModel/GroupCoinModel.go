package GroupCoinModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_coin"

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) Api_insert(group_id, user_id, cid, amount any) bool {
	db := self.Db.Table(table)
	data := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
		"cid":      cid,
		"amount":   amount,
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

func (self *Interface) Api_find(group_id, user_id, cid any) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
		"cid":      cid,
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

func (self *Interface) Api_select(group_id, user_id any) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
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

func (self *Interface) Api_sum_byCid(cid any) float64 {
	db := self.Db.Table(table)
	where := map[string]any{
		"cid": cid,
	}
	db.Where(where)
	ret, err := db.Sum("amount")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		if ret != nil {
			return ret.(float64)
		} else {
			return 0
		}
	}
}

func (self *Interface) Api_join_select(group_id, user_id any) []gorose.Data {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	db.Join("coin on coin.id=cid")
	db.Where("amount", ">", 0)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_incr(group_id, user_id, cid, amount any) bool {
	db := self.Db.Table(table)
	where := map[string]any{
		"group_id": group_id,
		"user_id":  user_id,
		"cid":      cid,
	}
	db.Where(where)
	_, err := db.Increment("amount", amount)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
