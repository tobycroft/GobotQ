package GroupFunction

import (
	"errors"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/action/GroupBalanceAction"
	"main.go/app/bot/action/MessageBuilder"
	"main.go/app/bot/model/DaojuModel"
	"main.go/app/bot/model/GroupDaojuModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/service"
	"main.go/config/app_default"
	"main.go/tuuz"

	"math"
	"strings"
)

func App_group_daoju(self_id, group_id, user_id, message_id int64, message string, groupmember map[string]any, groupfunction map[string]any) {
	switch message {
	case "清空我的背包":
		str := clear_backpack(group_id, user_id)
		AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(str), groupfunction)
		break

	case "我的", "列表", "背包":
		str := list_my_daoju(group_id, user_id)
		AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(str), groupfunction)
		break

	case "商城", "商店":
		str := list_daoju()
		AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(str), groupfunction)
		break

	case "帮助":
		AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(app_default.Default_daoju), groupfunction)
		break

	case "赠送":
		AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(app_default.Default_send_daoju), groupfunction)
		break

	case "购买", "兑换":
		AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(app_default.Daoju_goumai), groupfunction)
		break

	default:
		str, has := service.Serv_text_match(message, []string{"购买", "兑换"})
		if has {
			str, err := buy_daoju(group_id, user_id, str)
			if err != nil {
				AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(err.Error()), groupfunction)
			} else {
				AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(str), groupfunction)
			}
		}
		send_to, has := service.Serv_text_match(message, []string{"赠送", "赠与", "送给"})
		if has {
			send, err := send_daoju(group_id, user_id, send_to)
			if err != nil {
				AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(err.Error()), groupfunction)
			} else {
				AutoMessage(self_id, group_id, user_id, MessageBuilder.IMessageBuilder{}.New().Text(send), groupfunction)
			}
		}
		break
	}
}

func list_daoju() string {
	str := ""
	str += "目前可兑换的道具有："
	datas := DaojuModel.Api_select_canShow()
	for i, data := range datas {
		list := i + 1
		str += "\r\n	" + Calc.Int2String(list) + "." + data["cname"].(string) + "：" + Calc.Any2String(data["price"]) + "威望," + data["info"].(string)
	}
	str += "\r\n你可以使用“道具兑换”[道具名称]，例如“道具购买免死金牌”来购买对应的道具，或者使用“acfur道具”来查看帮助"
	return str
}

func buy_daoju(group_id, user_id, cname any) (string, error) {
	data := DaojuModel.Api_find_byCname(cname)
	if len(data) < 1 {
		return "", errors.New(app_default.Daoju_notfound)
	}
	db := tuuz.Db()
	db.Begin()

	var gbal GroupBalanceAction.Interface
	gbal.Db = db
	err, left := gbal.App_single_balance(group_id, user_id, nil, -math.Abs(data["price"].(float64)), "购买道具")
	if err != nil {
		db.Rollback()
		return "", err
	}
	var dj GroupDaojuModel.Interface
	dj.Db = db
	user_daoju_data := dj.Api_find(group_id, user_id, data["id"])
	if len(user_daoju_data) > 0 {
		if !dj.Api_incr(group_id, user_id, data["id"], 1) {
			db.Rollback()
			return "", errors.New("购买道具失败")
		}
	} else {
		if !dj.Api_insert(group_id, user_id, data["id"], 1) {
			db.Rollback()
			return "", errors.New("购买道具失败")
		}
	}
	ujd := dj.Api_find(group_id, user_id, data["id"])
	db.Commit()
	str := "您当前还剩" + Calc.Any2String(left) + "威望\r\n"
	str += "您当前拥有" + Calc.Any2String(ujd["num"]) + "个同类型道具"
	return "兑换完成，您兑换了：" + Calc.Any2String(data["cname"]) + "" +
		"\r\n " + Calc.Any2String(data["cname"]) + ":" + Calc.Any2String(data["info"]) +
		"\r\n 类型:" + Calc.Any2String(data["type"]) + "\r\n 消耗:" + Calc.Any2String(data["price"]) +
		"\r\n" + str, nil
}

func clear_backpack(group_id, user_id any) string {
	datas := GroupDaojuModel.Api_select(group_id, user_id)
	str := "您已经清空了您的背包，如下道具被丢弃："
	for i, data := range datas {
		list := i + 1
		daoju := DaojuModel.Api_find(data["dj_id"])
		if len(daoju) < 1 {
			continue
		}
		str += "\r\n	" + Calc.Int2String(list) + "." + daoju["cname"].(string) + ",数量," + Calc.Any2String(data["num"])
	}
	var gjd GroupDaojuModel.Interface
	gjd.Db = tuuz.Db()
	if gjd.Api_delete(group_id, user_id) {
		return str
	} else {
		return "背包清空失败"
	}
}

func list_my_daoju(group_id, user_id any) string {
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

func send_daoju(group_id, user_id any, send_to_message string) (string, error) {
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
		db.Begin()
		var gd GroupDaojuModel.Interface
		gd.Db = db
		user_daoju := gd.Api_value_num(group_id, user_id, daoju_data["id"])
		if user_daoju < 1 {
			db.Rollback()
			return "", errors.New("你没有这个道具，无法赠送")
		} else {
			if !gd.Api_decr(group_id, user_id, daoju_data["id"]) {
				db.Rollback()
				return "", errors.New("赠送失败，无法扣除该道具")
			}
			if len(gd.Api_find(group_id, member["user_id"], daoju_data["id"])) < 1 {
				if !gd.Api_insert(group_id, member["user_id"], daoju_data["id"], 1) {
					db.Rollback()
					return "", errors.New("赠送失败，对方无法增加该类型道具")
				}
			} else {
				if !gd.Api_incr(group_id, member["user_id"], daoju_data["id"], 1) {
					db.Rollback()
					return "", errors.New("赠送失败，对方无法增加该类型道具")
				}
			}
			left := gd.Api_value_num(group_id, user_id, daoju_data["id"])
			db.Commit()
			return "成功赠送道具，你目前还剩" + Calc.Any2String(left) + "个" + send_to_message, nil
		}
	} else {
		return "", errors.New("该道具不存在")
	}
}
