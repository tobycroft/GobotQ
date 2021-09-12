package Group

import (
	"errors"
	"main.go/app/bot/api"
	"main.go/app/bot/model/DaojuModel"
	"main.go/app/bot/model/GroupBalanceModel"
	"main.go/app/bot/model/GroupDaojuModel"
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
		played_time := GroupLunpanModel.Api_count(group_id)
		if played_time > 85 {
			played_time = 85
		}
		ext_text := ""
		if active {
			possible := int64(0)
			var gd GroupDaojuModel.Interface
			gd.Db = db
			user_daoju := gd.Api_find_in_djId(group_id, user_id, []interface{}{4, 5, 6, 7})
			if len(user_daoju) > 0 {
				daoju := DaojuModel.Api_find_canUse(user_daoju["dj_id"])
				if len(daoju) > 0 {
					switch daoju["name"].(string) {
					case "r_25":
						possible = 25
						break

					case "r_50":
						possible = 50
						break

					case "r_75":
						possible = 75
						break

					case "r_90":
						possible = 90
						break

					default:
						break
					}
				}
			}

			if possible > played_time {
				played_time = possible
				ext_text = ",你用了自家的子弹，这课弹的激发概率为:" + Calc.Any2String(100-played_time) + "％"
			} else {
				ext_text = ",左轮目前完好度:" + Calc.Any2String(100-played_time) + "％"
			}
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

			rand_num := Calc.Rand(10, 60)
			rand := Calc.Float642String(math.Floor(float64(rand_num / 10)))
			rand_slice := []string{}
			stuck_mode := int64(Calc.Rand(1, 100))

			daoju_num := gd.Api_value_num(group_id, user_id, 3)
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
				poom := Array.InArrayString(rand, rand_slice)
				if poom && stuck_mode > played_time {
					//poom!!!
					if daoju_num > 0 {
						amount = 0
						gd.Api_decr(group_id, user_id, 3)
						AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n-Dang!\n脖子差点折了，子弹被放在了位置"+tick+"上，"+
							"激发位置在"+rand+",因为你带了防弹头盔，所以平局，不奖励也不损失威望"+ext_text, groupfunction)
					} else {
						amount = -num
						AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n可惜了，子弹被放在了位置"+tick+"上，"+
							"激发位置在"+rand+",因此你损失了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
					}
				} else if poom && stuck_mode <= played_time {
					amount = Calc.Round(num/6, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick~\n好险！子弹被放在了位置"+tick+"上，"+
						"激发位置在"+rand+",卡弹了，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
				} else {
					amount = Calc.Round(num/6, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Kaa~\nWow赢了！子弹被放在了位置"+tick+"上，"+
						"激发位置在"+rand+",没响，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
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
				poom := Array.InArrayString(rand, rand_slice)
				if poom && stuck_mode > played_time {
					//poom!!!
					if daoju_num > 0 {
						amount = 0
						gd.Api_decr(group_id, user_id, 3)
						AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n-Dang!\n脖子差点折了，子弹被放在了位置"+tick+"上，"+
							"激发位置在"+rand+",因为你带了防弹头盔，所以平局，不奖励也不损失威望"+ext_text, groupfunction)
					} else {
						amount = -num
						AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n可惜了，子弹被放在了位置"+tick+"上，"+
							"激发位置在"+rand+",因此你损失了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
					}
				} else if poom && stuck_mode <= played_time {
					amount = Calc.Round(num/3, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick~\n好险！子弹被放在了位置"+tick+"上，"+
						"激发位置在"+rand+",卡弹了，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
				} else {
					amount = Calc.Round(num/3, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Kaa~\nWow赢了！子弹被放在了位置"+tick+"上，"+
						"激发位置在"+rand+",没响，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
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
				poom := Array.InArrayString(rand, rand_slice)
				if poom && stuck_mode > played_time {
					//poom!!!
					if daoju_num > 0 {
						amount = 0
						gd.Api_decr(group_id, user_id, 3)
						AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n-Dang!\n脖子差点折了，子弹被放在了位置"+tick+"上，"+
							"激发位置在"+rand+",因为你带了防弹头盔，所以平局，不奖励也不损失威望"+ext_text, groupfunction)
					} else {
						amount = -num
						AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n可惜了，子弹被放在了位置"+tick+"上，"+
							"激发位置在"+rand+",因此你损失了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
					}
				} else if poom && stuck_mode <= played_time {
					amount = Calc.Round(num/2, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick~\n好险！子弹被放在了位置"+tick+"上，"+
						"激发位置在"+rand+",卡弹了，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
				} else {
					amount = Calc.Round(num/2, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Kaa~\nWow赢了！子弹被放在了位置"+tick+"上，"+
						"激发位置在"+rand+",没响，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
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
				poom := Array.InArrayString(rand, rand_slice)
				if poom && stuck_mode > played_time {
					//poom!!!
					if daoju_num > 0 {
						amount = 0
						gd.Api_decr(group_id, user_id, 3)
						AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n-Dang!\n脖子差点折了，子弹被放在了位置"+tick+"上，"+
							"激发位置在"+rand+",因为你带了防弹头盔，所以平局，不奖励也不损失威望"+ext_text, groupfunction)
					} else {
						amount = -num
						AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n可惜了，子弹被放在了位置"+tick+"上，"+
							"激发位置在"+rand+",因此你损失了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
					}
				} else if poom && stuck_mode <= played_time {
					amount = Calc.Round(num/3*2, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick~\n好险！子弹被放在了位置"+tick+"上，"+
						"激发位置在"+rand+",卡弹了，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
				} else {
					amount = Calc.Round(num/3*2, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Kaa~\nWow赢了！子弹被放在了位置"+tick+"上，"+
						"激发位置在"+rand+",没响，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
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
				poom := Array.InArrayString(rand, rand_slice)
				if poom && stuck_mode > played_time {
					//poom!!!
					if daoju_num > 0 {
						amount = 0
						gd.Api_decr(group_id, user_id, 3)
						AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n-Dang!\n脖子差点折了，子弹被放在了位置"+tick+"上，"+
							"激发位置在"+rand+",因为你带了防弹头盔，所以平局，不奖励也不损失威望"+ext_text, groupfunction)
					} else {
						amount = -num
						AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n可惜了，子弹被放在了位置"+tick+"上，"+
							"激发位置在"+rand+",因此你损失了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
					}
				} else if poom && stuck_mode <= played_time {
					amount = Calc.Round(num/6*5, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick~\n好险！子弹被放在了位置"+tick+"上，"+
						"激发位置在"+rand+",卡弹了，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
				} else {
					amount = Calc.Round(num/6*5, 2)
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Kaa~\nWow赢了！子弹被放在了位置"+tick+"上，"+
						"激发位置在"+rand+",没响，你成功得到了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
				}
				break

			case "F", "f":
				//poom!!!
				if stuck_mode > played_time {
					if daoju_num > 0 {
						amount = 0
						gd.Api_decr(group_id, user_id, 3)
						AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n-Dang!\n脖子差点折了，因为你带了防弹头盔，所以平局，不奖励也不损失威望"+ext_text, groupfunction)
					} else {
						amount = -num
						AutoMessage(self_id, group_id, user_id, at+"\n-Tick!\n-Poom！\n必死结局，你白白损失了"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
					}
				} else {
					amount = num
					AutoMessage(self_id, group_id, user_id, at+"\n-Tick~\n百分之20的卡弹率让你碰上了！恭喜你！运气爆棚奖励翻倍，你赢得了:"+Calc.Any2String(math.Abs(amount))+"威望~"+ext_text, groupfunction)
				}
				break

			default:
				AutoMessage(self_id, group_id, user_id, at+"请输入一个正确的字母，想参与1/6胜率轮盘输入“轮盘A10”，2/6输入“轮盘B10”，3/6选C，以此类推可在ABCDE中选择(大小写不敏感)"+ext_text, groupfunction)
				return
			}
			if !gbp.Api_incr(group_id, user_id, amount) {
				db.Rollback()
				Log.Errs(errors.New("GroupBalanceModel,增加失败"), tuuz.FUNCTION_ALL())
				return
			}
		} else {
			ext_text = ",左轮目前完好度:" + Calc.Any2String(100-played_time) + "％"
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
			str += ext_text
			count_lunpan := GroupLunpanModel.Api_count_userId(group_id, user_id)
			if count_lunpan == 0 {
				str += "\n\n这是你第一次参与轮盘，下次你可以用“轮盘[模式字母][数字]" +
					"\n例如“轮盘A10”，来挑战1/6的失败几率，挑战成功，奖励1/6押注威望" +
					"\n同时你可以使用轮盘B10来挑战2/6的败率，获得2/6的奖励" +
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
