package BalanceAction

import (
	"errors"
	"github.com/gohouse/gorose/v2"
	"main.go/app/v1/user/model/UserBalanceModel"
	"main.go/tuuz"
)

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) App_user_balance(uid interface{}) error {
	userbalance := self.App_check_balance(uid)

}

func (self *Interface) App_check_balance(uid interface{}) (float64, error) {
	self_create := false
	if self.Db == nil {
		self.Db = tuuz.Db()
		self.Db.Begin()
		self_create = true
	}
	var ub UserBalanceModel.Interface
	ub.Db = self.Db
	userbalance := ub.Api_find(uid)
	if len(userbalance) > 0 {
		if self_create {
			self.Db.Commit()
		}
		return userbalance["balance"].(float64), nil
	} else {
		if ub.Api_insert(uid, 0) {
			if self_create {
				self.Db.Commit()
			}
			return 0, nil
		} else {
			if self_create {
				self.Db.Rollback()
			}
			return 0, errors.New("UB初始化失败")
		}
	}
}
