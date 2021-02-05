package BalanceAction

import (
	"github.com/gohouse/gorose/v2"
	"main.go/app/v1/user/model/UserBalanceModel"
)

type Interface struct {
	Db gorose.IOrm
}

func (self *Interface) App_user_balance(uid interface{}) error {
	userbalance := UserBalanceModel.Api_find(uid)
	if len(userbalance)
}

func (self *Interface) App_check_balance(uid interface{}) (float64, error) {
	var ub UserBalanceModel.Interface
	ub.Db = self.Db
	userbalance := ub.Api_find(uid)
	if len(userbalance)>0{
		ub.Api_dec_balance()
	}else{

	}
}
