package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/FriendListModel"
)

func App_fresh_group_list() {
	bots := BotModel.Api_select()
	for _, bot := range bots {
		fl, err := api.Getgrouplist(bot["bot"])
		if err != nil {

		} else {
			FriendListModel.Api_delete(bot["bot"])
			var fss []FriendListModel.FriendList
			for _, fll := range fl {
				var fs FriendListModel.FriendList
				fs.Bot = bot["bot"]
				fs.Uid = fll.UIN
				fs.Nickname = fll.NickName
				fs.Email = fll.Email
				fs.Remark = fll.Remark
				fss = append(fss, fs)
			}
			FriendListModel.Api_insert_more(fss)
		}
	}
}
