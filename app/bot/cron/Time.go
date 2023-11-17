package cron

//func BaseCron() {
//	for {
//		bots := BotModel.Api_select()
//		for _, bot := range bots {
//			self_id := bot["self_id"]
//			gl, err := iapi.Api.Getgrouplist(bot["self_id"])
//			if err != nil {
//
//			} else {
//
//				GroupListRedis.Cac_del(self_id, "*")
//				GroupListModel.Api_delete(self_id)
//				GroupAdminModel.Api_delete_bySelfIdAndGroupId(self_id, nil)
//				var gss []GroupListModel.GroupList
//				var gas []GroupAdminModel.GroupAdmins
//				for _, gll := range gl {
//					var gs GroupListModel.GroupList
//					gs.SelfId = self_id
//					gs.GroupId = gll.GroupId
//					gs.GroupName = gll.GroupName
//					gs.GroupMemo = gll.GroupRemark
//					gs.MemberCount = gll.MemberCount
//					gs.MaxMemberCount = gll.MaxMemberCount
//					gs.Admins, _ = Jsong.Encode(gll.Admins)
//					gss = append(gss, gs)
//					for _, admin := range gll.Admins {
//						gas = append(gas, GroupAdminModel.GroupAdmins{
//							SelfId:  self_id,
//							GroupId: gll.GroupId,
//							UserId:  admin,
//						})
//					}
//				}
//				if len(gss) > 0 && len(gas) > 0 {
//					GroupListModel.Api_insert_more(gss)
//					GroupAdminModel.Api_insert_more(gas)
//				}
//			}
//		}
//		time.Sleep(3600 * time.Second)
//	}
//}
//func BotInfoCron() {
//	for {
//		bots := BotModel.Api_select()
//		for _, bot := range bots {
//			bot_info, err := iapi.Api.GetLoginInfo(bot["self_id"])
//			if err != nil {
//
//			} else {
//				self_id := bot_info.UserID
//				name := bot_info.Nickname
//				if !BotModel.Api_update_cname(self_id, name) {
//					Log.Crrs(errors.New("机器人用户名无法更新"), tuuz.FUNCTION_ALL())
//				} else {
//					fmt.Println("机器人更新：", bot_info)
//				}
//			}
//		}
//		time.Sleep(30 * time.Minute)
//	}
//}
