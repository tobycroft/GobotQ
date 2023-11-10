package GroupBalanceAction

import (
	"errors"
	"github.com/tobycroft/Calc"
	"github.com/tobycroft/gorose-pro"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

type Interface struct {
	Db gorose.IOrm
}

func App_single_balance(group_id, user_id interface{}, order_id interface{}, amount float64, remark string) (error, float64) {
	var self Interface
	self.Db = tuuz.Db()
	return self.App_single_balance(group_id, user_id, order_id, amount, remark)
}

func (self *Interface) App_single_balance(group_id, user_id interface{}, order_id interface{}, amount float64, remark string) (error, float64) {
	self.Db.Begin()
	userbalance, err := self.App_check_balance(group_id, user_id)
	if err != nil {
		self.Db.Rollback()
		Log.Crrs(err, tuuz.FUNCTION_ALL())
		return err, 0
	}
	if order_id == nil {
		order_id = Calc.GenerateOrderId()
	}
	after_userbalance, _ := Calc.Bc_add(userbalance, amount).Float64()
	if amount < 0 {
		if after_userbalance < 0 {
			self.Db.Rollback()
			return errors.New("威望不足无法购买,还差：" + Calc.Float642String(after_userbalance)), 0
		}
	}
	var ub GroupBalanceModel.Interface
	ub.Db = self.Db
	if !ub.Api_incr(group_id, user_id, amount) {
		self.Db.Rollback()
		return errors.New("威望变动出现故障"), 0
	}
	self.Db.Commit()
	return nil, after_userbalance
}

func (self *Interface) App_check_balance(group_id, user_id interface{}) (float64, error) {
	var ub GroupBalanceModel.Interface
	ub.Db = self.Db
	userbalance := ub.Api_find(group_id, user_id)
	if len(userbalance) > 0 {
		return userbalance["balance"].(float64), nil
	} else {
		if ub.Api_insert(group_id, user_id) {
			return 0, nil
		} else {
			return 0, errors.New("威望初始化失败")
		}
	}
}
