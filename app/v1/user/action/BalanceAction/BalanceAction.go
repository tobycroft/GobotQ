package BalanceAction

import (
	"errors"
	"github.com/tobycroft/gorose"
	"main.go/app/v1/user/model/UserBalanceModel"
	"main.go/app/v1/user/model/UserBalanceRecordModel"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
)

type Interface struct {
	Db gorose.IOrm
}

func App_single_balance(uid interface{}, order_id interface{}, amount float64, remark string) error {
	var self Interface
	self.Db = tuuz.Db()
	return self.App_single_balance(uid, order_id, amount, remark)
}

func (self *Interface) App_single_balance(uid interface{}, order_id interface{}, amount float64, remark string) error {
	self.Db.Begin()
	userbalance, err := self.App_check_balance(uid)
	if err != nil {
		Log.Crrs(err, tuuz.FUNCTION_ALL())
		return err
	}
	if order_id == nil {
		order_id = Calc.GenerateOrderId()
	}
	after_userbalance, _ := Calc.Bc_add(userbalance, amount).Float64()
	if after_userbalance < 0 {
		self.Db.Rollback()
		return errors.New("余额还差" + Calc.Float642String(after_userbalance))
	}
	var ub UserBalanceModel.Interface
	ub.Db = self.Db
	if !ub.Api_inc_balance(uid, amount) {
		self.Db.Rollback()
		return errors.New("余额变动出现故障")
	}
	var ubr UserBalanceRecordModel.Interface
	ubr.Db = self.Db
	one_balancerecord := ubr.Api_find(uid)
	if len(one_balancerecord) > 0 {
		after_balancerecord, _ := Calc.Bc_add(one_balancerecord["after_balance"], amount).Float64()
		if after_userbalance < 0 {
			self.Db.Rollback()
			return errors.New("余额不足,仍需" + Calc.Float642String(after_balancerecord))
		}
		if !ubr.Api_insert(uid, one_balancerecord["after_balance"], amount, after_balancerecord, remark) {
			self.Db.Rollback()
			return errors.New("UserBalanceRecordModel插入失败")
		}
	} else {
		if !ubr.Api_insert(uid, 0, amount, amount, remark) {
			self.Db.Rollback()
			return errors.New("UserBalanceRecordModel插入失败")
		}
	}
	self.Db.Commit()
	return nil
}

func (self *Interface) App_check_balance(uid interface{}) (float64, error) {
	var ub UserBalanceModel.Interface
	self.Db.Begin()
	ub.Db = self.Db
	userbalance := ub.Api_find(uid)
	if len(userbalance) > 0 {
		self.Db.Commit()
		return userbalance["balance"].(float64), nil
	} else {
		if ub.Api_insert(uid, 0) {
			self.Db.Commit()
			return 0, nil
		} else {
			self.Db.Rollback()
			return 0, errors.New("UB初始化失败")
		}
	}
}
