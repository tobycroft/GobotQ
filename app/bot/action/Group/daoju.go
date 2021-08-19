package Group

import (
	"errors"
	"main.go/app/bot/action/GroupBalance"
	"main.go/app/bot/model/DaojuModel"
	"main.go/app/bot/model/GroupDaojuModel"
	"main.go/tuuz"
	"main.go/tuuz/Calc"
)

func App_group_daoju(self_id, group_id, user_id, message_id int64, message string, groupmember map[string]interface{}, groupfunction map[string]interface{}) {

}

func list_daoju() string {
	str := ""
	str += "目前可兑换的道具有："
	datas := DaojuModel.Api_select_canShow()
	for i, data := range datas {
		list := i + 1
		str += "\r\n" + Calc.Int2String(list) + "." + data["cname"].(string) + "：" + Calc.Any2String(data["price"]) + "威望," + data["info"].(string)
	}
	str += "\r\n你可以使用“道具兑换”[道具名称]，例如“道具购买免死金牌”来购买对应的道具，或者使用“acfur道具”来查看帮助"
	return str
}

func buy_daoju(group_id, user_id, cname interface{}) error {
	data := DaojuModel.Api_find_byCname(cname)
	db := tuuz.Db()
	db.Begin()

	var gbal GroupBalance.Interface
	gbal.Db = db
	err := gbal.App_single_balance(group_id, user_id, nil, data["price"].(float64), "购买道具")
	if err != nil {
		db.Rollback()
		return err
	}
	var dj GroupDaojuModel.Interface
	dj.Db = db
	user_daoju_data := dj.Api_find(group_id, user_id, data["id"])
	if len(user_daoju_data) > 0 {
		if dj.Api_incr(group_id, user_id, data["id"], 1) {
			return nil
		} else {
			db.Rollback()
			return errors.New("购买道具失败")
		}
	} else {
		if dj.Api_insert(group_id, user_id, data["id"], 1) {
			return nil
		} else {
			db.Rollback()
			return errors.New("购买道具失败")
		}
	}
}

func clean_backpack(group_id, user_id interface{}) string {
	datas := GroupDaojuModel.Api_select(group_id, user_id)
	str := "您已经清空了您的背包，如下道具被丢弃："
	for i, data := range datas {
		list := i + 1
		daoju := DaojuModel.Api_find(data["jd_id"])
		if len(daoju) < 1 {
			continue
		}
		str += "\r\n" + Calc.Int2String(list) + "." + daoju["cname"].(string) + ",数量," + Calc.Any2String(data["num"])
	}
	var gjd GroupDaojuModel.Interface
	gjd.Db = tuuz.Db()
	if gjd.Api_delete(group_id, user_id) {
		return str
	} else {
		return "背包清空失败"
	}
}

func send_daoju(group_id, user_id, to_uid interface{}) {

}
