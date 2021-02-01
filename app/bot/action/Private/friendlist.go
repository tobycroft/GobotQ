package Private

import (
	"main.go/app/bot/api"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/FriendListModel"
)

func App_refresh_friend_list_all() {
	bots := BotModel.Api_select()
	for _, bot := range bots {
		App_refresh_friend_list(bot["bot"])
	}
}

func App_refresh_friend_list(bot interface{}) {
	fl, err := api.Getfriendlist(bot)
	if err != nil {

	} else {
		FriendListModel.Api_delete(bot)
		var fss []FriendListModel.FriendList
		for _, fll := range fl {
			var fs FriendListModel.FriendList
			fs.Bot = bot
			fs.Uid = fll.UIN
			fs.Nickname = fll.NickName
			fs.Email = fll.Email
			fs.Remark = fll.Remark
			fss = append(fss, fs)
		}
		FriendListModel.Api_insert_more(fss)
	}
}
