package Group

import (
	"errors"
	"main.go/app/bot/action/GroupBalance"
	"main.go/app/bot/model/CoinModel"
	"main.go/app/bot/model/GroupCoinModel"
	"main.go/app/bot/service"
	"main.go/config/app_default"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
	"math"
	"strings"
)

func App_trade_center(self_id, group_id, user_id, message_id int64, message string, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	switch message {

	case "我的", "列表", "背包":
		str := list_my_coin(group_id, user_id)
		AutoMessage(self_id, group_id, user_id, str, groupfunction)
		break

	case "中心", "商城", "商店":
		str := list_coin()
		AutoMessage(self_id, group_id, user_id, str, groupfunction)
		break

	case "帮助":
		AutoMessage(self_id, group_id, user_id, app_default.Default_trade, groupfunction)
		break

	case "买入", "购买", "兑换":
		AutoMessage(self_id, group_id, user_id, app_default.Trade_buy, groupfunction)
		break
	case "卖出", "卖掉", "兑出":
		AutoMessage(self_id, group_id, user_id, app_default.Trade_sell, groupfunction)
		break

	default:
		str, has := service.Serv_text_match(message, []string{"买入", "购买", "兑换"})
		if has {
			str, err := buy_coin(group_id, user_id, str)
			if err != nil {
				AutoMessage(self_id, group_id, user_id, err.Error(), groupfunction)
			} else {
				AutoMessage(self_id, group_id, user_id, str, groupfunction)
			}
		}
		sell, has := service.Serv_text_match(message, []string{"卖出", "卖掉", "兑出"})
		if has {
			send, err := sell_coin(group_id, user_id, sell)
			if err != nil {
				AutoMessage(self_id, group_id, user_id, err.Error(), groupfunction)
			} else {
				AutoMessage(self_id, group_id, user_id, send, groupfunction)
			}
		}
		break
	}
}

func list_coin() string {
	str := ""
	str += "交易中心目前拥有如下币种可买入："
	datas := CoinModel.Api_select()
	for i, data := range datas {
		list := i + 1
		str += "\r\n	" + Calc.Int2String(list) + ".名称:" + data["cname"].(string) + ",比重:" + Calc.Any2String(data["price"]) +
			",增值率:" + Calc.Any2String(data["gain"]) + "(每天)" + ",说明:" + data["info"].(string)
	}
	str += "\r\n你可以使用“交易买入”[名称][数量]，例如“道具买入A100”来购买对应比重，例如比重为2时，使用100购买，就可以获得50比重，可使用“交易帮助”来查看帮助详情"
	return str
}

func buy_coin(group_id, user_id interface{}, message string) (string, error) {
	temp_amount := service.Serv_get_num(message)
	amount, err := Calc.Any2Float64_2(temp_amount)
	if err != nil {
		return "", errors.New(app_default.Trade_buy)
	}
	cname := strings.ReplaceAll(message, temp_amount, "")
	coin := CoinModel.Api_find_byCname(cname)
	if len(coin) < 1 {
		return "", errors.New(app_default.Trade_coin_not_found)
	}
	db := tuuz.Db()
	db.Begin()
	var gbal GroupBalance.Interface
	gbal.Db = db
	coin_num := amount / coin["price"].(float64)
	err, left := gbal.App_single_balance(group_id, user_id, nil, -math.Abs(amount), "购买币种")
	if err != nil {
		db.Rollback()
		return "", err
	}
	var gc GroupCoinModel.Interface
	gc.Db = db
	all_num := gc.Api_sum_byCid(coin["id"])
	user_coin := gc.Api_find(group_id, user_id, coin["id"])
	if len(user_coin) < 1 {
		if !gc.Api_insert(group_id, user_id, coin["id"], coin_num) {
			db.Rollback()
			return "", errors.New("买入记录创建失败")
		}
	} else {
		if !gc.Api_incr(group_id, user_id, coin["id"], coin_num) {
			db.Rollback()
			return "", errors.New("买入记录增加失败")
		}
	}
	switch coin["type"].(int64) {
	case 1:
		break

	case 2:
		after_all_num := gc.Api_sum_byCid(coin["id"])
		ratio := Calc.Round(all_num/after_all_num, 6)
		if ratio > 0 {
			var c CoinModel.Interface
			c.Db = db
			c.Api_incr_price(coin["id"], ratio)
		}
		break

	default:
		break
	}
	db.Commit()
	str := "您当前还剩" + Calc.Any2String(left) + "威望\r\n"
	return "您本次买入了" + Calc.Any2String(coin_num) + "个" + Calc.Any2String(coin["cname"]) +
		"\r\n 增长率:" + Calc.Any2String(coin["gain"]) + "\r\n 消耗:" + temp_amount +
		"\r\n" + str, nil
}

func sell_coin(group_id, user_id interface{}, message string) (string, error) {
	temp_amount := service.Serv_get_num(message)
	amount, err := Calc.Any2Float64_2(temp_amount)
	if err != nil {
		return "", errors.New(app_default.Trade_sell)
	}
	cname := strings.ReplaceAll(message, temp_amount, "")
	coin := CoinModel.Api_find_byCname(cname)
	if len(coin) < 1 {
		return "", errors.New(app_default.Trade_coin_not_found)
	}
	db := tuuz.Db()
	db.Begin()

	coin_reward := amount * coin["price"].(float64)
	var gc GroupCoinModel.Interface
	gc.Db = db
	user_coin := gc.Api_find(group_id, user_id, coin["id"])
	if len(user_coin) < 1 {
		db.Rollback()
		return "", errors.New("你没有这个币种，无法卖出")
	} else {
		if user_coin["amount"].(float64) < amount {
			db.Rollback()
			return "", errors.New("币数量不足，最多只能卖出:" + Calc.Any2String(user_coin["amount"]) + "个")
		}
		if !gc.Api_incr(group_id, user_id, coin["id"], -math.Abs(amount)) {
			db.Rollback()
			return "", errors.New("卖出记录减少失败")
		}
	}

	var gbal GroupBalance.Interface
	gbal.Db = db
	err, left := gbal.App_single_balance(group_id, user_id, nil, math.Abs(coin_reward), "卖出币种")
	if err != nil {
		db.Rollback()
		return "", err
	}
	db.Commit()
	str := "您当拥有" + Calc.Any2String(left) + "威望\r\n"
	return "您本次卖出了" + temp_amount + "个" + Calc.Any2String(coin["cname"]) +
		"\r\n 获得了" + Calc.Any2String(coin_reward) + "个威望\r\n" + str, nil
}

func list_my_coin(group_id, user_id interface{}) string {
	var gc GroupCoinModel.Interface
	gc.Db = tuuz.Db()
	datas := gc.Api_join_select(group_id, user_id)
	if len(datas) > 0 {
		str := "您拥有如下币种："
		for i, data := range datas {
			list := i + 1
			str += "\r\n	" + Calc.Int2String(list) + "." + data["cname"].(string) + ",数量:" + Calc.Any2String(data["amount"])
		}
		return str
	} else {
		return "您还未拥有任何币种，赶快购买一个来增加威望吧！"
	}
}
