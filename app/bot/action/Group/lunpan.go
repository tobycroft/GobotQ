package Group

import (
	"errors"
	"main.go/app/bot/api"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/app/bot/model/GroupLunpanModel"
	"main.go/app/bot/service"
	"main.go/tuuz"
	"main.go/tuuz/Array"
	"main.go/tuuz/Calc"
	"main.go/tuuz/Log"
	"math"
	"regexp"
	"strings"
)

func App_group_lunpan(self_id, group_id, user_id, message_id int64, message string, groupmember map[string]interface{}, groupfunction map[string]interface{}) {
	sign := GroupLunpanModel.Api_find(group_id, user_id)
	if groupfunction["sign_send_retract"].(int64) == 1 {
		var ret api.Struct_Retract
		ret.MessageId = message_id
		ret.Self_id = self_id
		api.Retract_chan <- ret
	}
	mode := regexp.MustCompile("[A-Za-z]")
	//fmt.Println(len(message), message, mode.MatchString(message))
	if len(message) > 3 && message[:3] != "" && !mode.MatchString(message) {
		//fmt.Println(len(message) > 3, message[:3] != "", !mode.MatchString(message))
		return
	}
	if len(sign) > 0 {
		at := service.Serv_at(user_id)
		AutoMessage(self_id, group_id, user_id, "你今天已经挑战过了，请明天再来"+at, groupfunction)
	} else {
		amount := float64(0)
		at := service.Serv_at(user_id)
		group_model := GroupBalanceModel.Api_find(group_id, user_id)
		rest_bal := float64(0)
		if group_model["balance"] == nil {
			rest_bal = 0
		} else {
			rest_bal = group_model["balance"].(float64)
		}

		if rest_bal < 0 {
			AutoMessage(self_id, group_id, user_id, at+"威望小于0,请先通过每日签到增加威望至正数", groupfunction)
			return
		}
		db := tuuz.Db()
		db.Begin()
		var gbp GroupBalanceModel.Interface
		gbp.Db = db
		if len(group_model) < 1 {
			if !gbp.Api_insert(group_id, user_id) {
				db.Rollback()
				Log.Errs(errors.New("GroupBalanceModel,写入失败"), tuuz.FUNCTION_ALL())
				return
			}
		}

		reg := regexp.MustCompile("[0-9]+")
		active := reg.MatchString(message)
		if active {
			//左轮模式
			mode_string := mode.FindString(message)
			message_num := reg.FindString(message)
			num, err := Calc.Any2Float64_2(message_num)
			if err != nil {
				AutoMessage(self_id, group_id, user_id, at+"想请输入一个正确的轮盘数字哦，不要超过自己的威望，可以使用[威望查询]来查看自己的威望", groupfunction)
				return
			}
			if num > rest_bal {
				AutoMessage(self_id, group_id, user_id, at+"你最多只能提取"+Calc.Any2String(rest_bal)+"威望参与游戏~", groupfunction)
				return
			}
			rand := Calc.Int2String(Calc.Rand(1, 6))
			rand_slice := []string{}
			switch mode_string {

			case "A", "a":
				for i := 0; i < 1; i++ {
					r1 := Calc.Int2String(Calc.Rand(1, 6))
					if !Array.InArrayString(r1, rand_slice) {
						rand_slice = append(rand_slice, r1)
					} else {
						i = i - 1
					}
				}
				tick := strings.Join(rand_slice, ",")
				if Array.InArrayString(rand, rand_slice) {
					//poom!!!
					amount = -num
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n可惜了，子弹被放在了位置"+rand+"上，"+
						"激发位置在"+tick+",因此你损失了"+Calc.Any2String(math.Abs(amount))+"威望~", groupfunction)
				} else {
					amount = Calc.Round(num/6, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Kaa~\nWow赢了！子弹被放在了位置"+rand+"上，"+
						"激发位置在"+tick+",没响，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~", groupfunction)
				}
				break

			case "B", "b":
				for i := 0; i < 2; i++ {
					r1 := Calc.Int2String(Calc.Rand(1, 6))
					if !Array.InArrayString(r1, rand_slice) {
						rand_slice = append(rand_slice, r1)
					} else {
						i = i - 1
					}
				}
				tick := strings.Join(rand_slice, ",")
				if Array.InArrayString(rand, rand_slice) {
					//poom!!!
					amount = -num
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n可惜了，子弹被放在了位置"+rand+"上，"+
						"激发位置在"+tick+",因此你损失了"+Calc.Any2String(math.Abs(amount))+"威望~", groupfunction)
				} else {
					amount = Calc.Round(num/3, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Kaa~\nWow赢了！子弹被放在了位置"+rand+"上，"+
						"激发位置在"+tick+",没响，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~", groupfunction)
				}
				break

			case "C", "c":
				for i := 0; i < 3; i++ {
					r1 := Calc.Int2String(Calc.Rand(1, 6))
					if !Array.InArrayString(r1, rand_slice) {
						rand_slice = append(rand_slice, r1)
					} else {
						i = i - 1
					}
				}
				tick := strings.Join(rand_slice, ",")
				if Array.InArrayString(rand, rand_slice) {
					//poom!!!
					amount = -num
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n可惜了，子弹被放在了位置"+rand+"上，"+
						"激发位置在"+tick+",因此你损失了"+Calc.Any2String(math.Abs(amount))+"威望~", groupfunction)
				} else {
					amount = Calc.Round(num/2, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Kaa~\nWow赢了！子弹被放在了位置"+rand+"上，"+
						"激发位置在"+tick+",没响，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~", groupfunction)
				}
				break

			case "D", "d":
				for i := 0; i < 4; i++ {
					r1 := Calc.Int2String(Calc.Rand(1, 6))
					if !Array.InArrayString(r1, rand_slice) {
						rand_slice = append(rand_slice, r1)
					} else {
						i = i - 1
					}
				}
				tick := strings.Join(rand_slice, ",")
				if Array.InArrayString(rand, rand_slice) {
					//poom!!!
					amount = -num
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n可惜了，子弹被放在了位置"+rand+"上，"+
						"激发位置在"+tick+",因此你损失了"+Calc.Any2String(math.Abs(amount))+"威望~", groupfunction)
				} else {
					amount = Calc.Round(num/3*2, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Kaa~\nWow赢了！子弹被放在了位置"+rand+"上，"+
						"激发位置在"+tick+",没响，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~", groupfunction)
				}
				break

			case "E", "e":
				for i := 0; i < 5; i++ {
					r1 := Calc.Int2String(Calc.Rand(1, 6))
					if !Array.InArrayString(r1, rand_slice) {
						rand_slice = append(rand_slice, r1)
					} else {
						i = i - 1
					}
				}
				tick := strings.Join(rand_slice, ",")
				if Array.InArrayString(rand, rand_slice) {
					//poom!!!
					amount = -num
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n可惜了，子弹被放在了位置"+rand+"上，"+
						"激发位置在"+tick+",因此你损失了"+Calc.Any2String(math.Abs(amount))+"威望~", groupfunction)
				} else {
					amount = Calc.Round(num/6*5, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Kaa~\nWow赢了！子弹被放在了位置"+rand+"上，"+
						"激发位置在"+tick+",没响，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~", groupfunction)
				}
				break

			case "F", "f":
				//poom!!!
				amount = -num
				AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n必死结局，你白白损失了"+Calc.Any2String(math.Abs(amount))+"威望~", groupfunction)
				break

			default:
				AutoMessage(self_id, group_id, user_id, at+"请输入一个正确的字母，想参与1/6胜率轮盘输入A，2/6输入B，3/6选C，以此类推可在ABCDE中选择(大小写不敏感)", groupfunction)
				return
			}
			if !gbp.Api_incr(group_id, user_id, amount) {
				db.Rollback()
				Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
				return
			}
		} else {
			//普通模式
			rand := Calc.Rand(0, 100)
			str := ""
			if rand <= 1 {
				amount = Calc.Round(rest_bal*9, 2)
				str += at + "十倍奖励完胜,当前余额:" + Calc.Any2String(rest_bal+amount)
			} else if rand > 1 && rand <= 20 {
				amount = -float64(rand)
				str += at + "小败,扣除:" + Calc.Any2String(math.Abs(amount)) + ",当前余额:" + Calc.Any2String(rest_bal+amount)
			} else if rand > 20 && rand <= 50 {
				amount = 2
				str += at + "小胜,赢得2" + ",当前余额:" + Calc.Any2String(rest_bal+amount)
			} else if rand > 50 && rand <= 85 {
				amount = 5
				str += at + "胜利,赢得5" + ",当前余额:" + Calc.Any2String(rest_bal+amount)
			} else if rand > 85 && rand <= 95 {
				amount = 10
				str += at + "大胜,赢得10" + ",当前余额:" + Calc.Any2String(rest_bal+amount)
			} else if rand > 95 && rand <= 99 {
				amount = -Calc.Round(rest_bal/2, 2)
				str += at + "扣除一半轮盘大败,当前余额:" + Calc.Any2String(rest_bal+amount)
			} else {
				amount = -rest_bal
				str += at + "轮盘完败,你的余额已不复存在"
			}
			count_lunpan := GroupLunpanModel.Api_count_userId(group_id, user_id)
			if count_lunpan == 0 {
				str += "\n\n这是你第一次参与轮盘，下次你可以用“轮盘[模式字母][数字]" +
					"\n例如“轮盘A10”，来挑战1/6的获胜几率，挑战成功，奖励1/6押注威望" +
					"\n同时你可以使用轮盘B10来挑战2/6的胜率，获得2/6的奖励" +
					"\n可选模式有ABCDE，挑战威望无上限，你可以使用威望查询来查看自己的可用情况" +
					"\n觉得自己运气还不错的话可以试试哦~"
			}
			if !gbp.Api_incr(group_id, user_id, amount) {
				db.Rollback()
				Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
				return
			}
			AutoMessage(self_id, group_id, user_id, str, groupfunction)
		}

		var gsp GroupLunpanModel.Interface
		gsp.Db = db
		if !gsp.Api_insert(group_id, user_id) {
			db.Rollback()
			Log.Errs(errors.New("GroupLunpanModel,插入失败"), tuuz.FUNCTION_ALL())
			return
		} else {
			db.Commit()
		}
	}
}
