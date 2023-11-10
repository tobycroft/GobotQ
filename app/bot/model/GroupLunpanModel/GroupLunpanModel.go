package GroupLunpanModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_lunpan"

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
	db.LockForUpdate()
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func (self *Interface) Api_find(group_id, user_id interface{}) gorose.Data {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	db.Where("date > (SELECT CURDATE())")
	db.LockForUpdate()
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_count(group_id interface{}) int64 {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
	}
	db.Where(where)
	db.Where("date > (SELECT CURDATE())")
	db.LockForUpdate()
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func Api_select(group_id interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
	}
	db.Where(where)
	db.Where("date > (SELECT CURDATE())")
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func (self *Interface) Api_count_userId(group_id, user_id interface{}) int64 {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	db.LockForUpdate()
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}

func (self *Interface) Api_count_userIdandDate(group_id, user_id, date interface{}) int64 {
	db := self.Db.Table(table)
	where := map[string]interface{}{
		"group_id": group_id,
		"user_id":  user_id,
	}
	db.Where(where)
	db.Where("date", ">", date)
	db.LockForUpdate()
	ret, err := db.Count()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return 0
	} else {
		return ret
	}
}
