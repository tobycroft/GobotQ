package GroupSignModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Date"
	"main.go/tuuz/Log"
)

const table = "group_sign"

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

func Api_find(gid, uid interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"gid": gid,
		"uid": uid,
	}
	db.Where(where)
	db.Where("date", ">", Date.Today())
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_count(gid interface{}) int64 {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"gid": gid,
	}
	db.Where(where)
	db.Where("date", ">", Date.Today())
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}

}

func Api_select(gid interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"gid": gid,
	}
	db.Where(where)
	db.Where("date", ">", Date.Today())
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
