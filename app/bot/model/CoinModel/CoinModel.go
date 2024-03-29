package CoinModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "coin"

func Api_find(id any) gorose.Data {
	db := tuuz.Db().Table(table)
	db.Where("id", "=", id)
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

func Api_find_byCname(cname any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"cname": cname,
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

func Api_find_byName(name any) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"name": name,
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

func (self *Interface) Api_incr_price(id, price any) bool {
	db := self.Db.Table(table)
	db.Where("id", "=", id)
	_, err := db.Increment("price", price)
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
