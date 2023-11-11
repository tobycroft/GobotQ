package Private

import (
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/FriendListModel"
)

func App_refresh_friend_list_all() {
	bots := BotModel.Api_select()
	for _, bot := range bots {
		App_refresh_friend_list(bot["self_id"])
	}
}

func App_refresh_friend_list(self_id any) {
	fl, err := iapi.Post{}.Getfriendlist(self_id)
	if err != nil {

	} else {
		FriendListModel.Api_delete(self_id)
		var fss []FriendListModel.FriendList
		for _, fll := range fl {
			var fs FriendListModel.FriendList
			fs.SelfId = self_id
			fs.UserId = fll.UserID
			fs.Nickname = fll.Nickname
			fs.Remark = fll.Remark
			fss = append(fss, fs)
		}
		FriendListModel.Api_insert_more(fss)
	}
}
