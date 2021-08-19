package GroupBalance

import (
	"errors"
	"github.com/gohouse/gorose/v2"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
)

type Interface struct {
	Db gorose.IOrm
}

func App_single_balance(group_id, user_id interface{}, order_id interface{}, amount float64, remark string) error {
	var self Interface
	self.Db = tuuz.Db()
	return self.App_single_balance(group_id, user_id, order_id, amount, remark)
}

func (self *Interface) App_single_balance(group_id, user_id interface{}, order_id interface{}, amount float64, remark string) error {
	self.Db.Begin()
	userbalance, err := self.App_check_balance(group_id, user_id)
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
		return errors.New("威望还差：" + Calc.Float642String(after_userbalance))
	}
	var ub GroupBalanceModel.Interface
	ub.Db = self.Db
	if !ub.Api_incr(group_id, user_id, amount) {
		self.Db.Rollback()
		return errors.New("威望变动出现故障")
	}
	self.Db.Commit()
	return nil
}

func (self *Interface) App_check_balance(group_id, user_id interface{}) (float64, error) {
	var ub GroupBalanceModel.Interface
	self.Db.Begin()
	ub.Db = self.Db
	userbalance := ub.Api_find(group_id, user_id)
	if len(userbalance) > 0 {
		self.Db.Commit()
		return userbalance["balance"].(float64), nil
	} else {
		if ub.Api_insert(group_id, user_id) {
			self.Db.Commit()
			return 0, nil
		} else {
			self.Db.Rollback()
			return 0, errors.New("威望初始化失败")
		}
	}
}
