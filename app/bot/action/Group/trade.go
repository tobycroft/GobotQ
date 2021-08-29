package Group

import (
	"errors"
	"github.com/shopspring/decimal"
	"main.go/app/bot/action/GroupBalance"
	"main.go/app/bot/model/CoinModel"
	"main.go/app/bot/model/DaojuModel"
	"main.go/app/bot/model/GroupCoinModel"
	"main.go/app/bot/model/GroupDaojuModel"
	"main.go/app/bot/model/GroupMemberModel"
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
		str := list_my_daoju(group_id, user_id)
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
		send_to, has := service.Serv_text_match(message, []string{"卖出", "卖掉", "兑出"})
		if has {
			send, err := send_daoju(group_id, user_id, send_to)
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
	datas := DaojuModel.Api_select_canShow()
	for i, data := range datas {
		list := i + 1
		str += "\r\n	" + Calc.Int2String(list) + ".名称:" + data["cname"].(string) + ",比重:" + Calc.Any2String(data["price"]) + ",增值率:" + Calc.Any2String(data["gain"]) + "(每天)"
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

	price, _ := coin["price"].(decimal.Decimal).Abs().Float64()

	coin_num := amount / price

	err, left := gbal.App_single_balance(group_id, user_id, nil, -math.Abs(amount), "购买币种")
	if err != nil {
		db.Rollback()
		return "", err
	}
	var gc GroupCoinModel.Interface
	gc.Db = db
	user_coin := gc.Api_find(group_id, user_id, coin["id"])
	if len(user_coin) < 1 {
		if !gc.Api_insert(group_id, user_id, coin["id"], coin_num) {
			db.Rollback()
			return "", errors.New("买入记录创建失败")
		}
	} else {
		if !gc.Api_incr(group_id, user_id, coin["id"], coin_num) {
			db.Rollback()
			return "", errors.New("修改记录创建失败")
		}
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

	price, _ := coin["price"].(decimal.Decimal).Abs().Float64()

	coin_num := amount / price

	var gc GroupCoinModel.Interface
	gc.Db = db
	user_coin := gc.Api_find(group_id, user_id, coin["id"])
	if len(user_coin) < 1 {
		if !gc.Api_insert(group_id, user_id, coin["id"], coin_num) {
			db.Rollback()
			return "", errors.New("买入记录创建失败")
		}
	} else {
		if !gc.Api_incr(group_id, user_id, coin["id"], coin_num) {
			db.Rollback()
			return "", errors.New("修改记录创建失败")
		}
	}

	var gbal GroupBalance.Interface
	gbal.Db = db

	err, left := gbal.App_single_balance(group_id, user_id, nil, -math.Abs(amount), "购买币种")
	if err != nil {
		db.Rollback()
		return "", err
	}

	db.Commit()
	str := "您当前还剩" + Calc.Any2String(left) + "威望\r\n"
	return "您本次买入了" + Calc.Any2String(coin_num) + "个" + Calc.Any2String(coin["cname"]) +
		"\r\n 增长率:" + Calc.Any2String(coin["gain"]) + "\r\n 消耗:" + temp_amount +
		"\r\n" + str, nil
}

func list_my_daoju(group_id, user_id interface{}) string {
	datas := GroupDaojuModel.Api_select_have(group_id, user_id)
	if len(datas) > 0 {
		str := "您拥有如下道具："
		for i, data := range datas {
			list := i + 1
			daoju := DaojuModel.Api_find_canUse(data["dj_id"])
			if len(daoju) < 1 {
				continue
			}
			str += "\r\n	" + Calc.Int2String(list) + "." + daoju["cname"].(string) + ",作用" + Calc.Any2String(daoju["type"]) + ",数量," + Calc.Any2String(data["num"])
		}
		return str
	} else {
		return "您还未拥有任何道具,可以使用“道具商店”命令来查看可购买的道具"
	}
}

func send_daoju(group_id, user_id interface{}, send_to_message string) (string, error) {
	qq := service.Serv_get_qq(send_to_message)
	cq_mess, to_user_id := service.Serv_at_who(send_to_message)
	qq_num := ""
	if to_user_id != "" {
		qq_num = to_user_id
		send_to_message = strings.ReplaceAll(send_to_message, cq_mess, "")
	} else if qq != "" {
		qq_num = qq
		send_to_message = strings.ReplaceAll(send_to_message, qq, "")
	} else {
		return "", errors.New("接收人不正确，请使用道具赠送[道具名称]群成员号码或者道具赠送[道具名称]@某个人，例如“道具赠送免死金牌@张三”来赠送自己已有的道具")
	}
	member := GroupMemberModel.Api_find(group_id, qq_num)
	if len(member) < 1 {
		return "", errors.New("没有找到该群员")
	}
	daoju_data := DaojuModel.Api_find_byCname(send_to_message)
	if len(daoju_data) > 0 {
		db := tuuz.Db()
		var gd GroupDaojuModel.Interface
		gd.Db = db
		user_daoju := gd.Api_value_num(group_id, user_id, daoju_data["id"])
		if user_daoju < 1 {
			return "", errors.New("你没有这个道具，无法赠送")
		} else {
			db.Begin()
			if !gd.Api_decr(group_id, user_id, daoju_data["id"]) {
				db.Rollback()
				return "", errors.New("赠送失败，无法扣除该道具")
			}
			if !gd.Api_incr(group_id, member["user_id"], daoju_data["id"], 1) {
				db.Rollback()
				return "", errors.New("赠送失败，对方无法增加该类型道具")
			}
			left := gd.Api_value_num(group_id, user_id, daoju_data["id"])
			db.Commit()
			return "成功赠送道具，你目前还剩" + Calc.Any2String(left) + "个" + send_to_message, nil
		}
	} else {
		return "", errors.New("该道具不存在")
	}
}
