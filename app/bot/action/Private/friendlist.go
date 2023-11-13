package Private

import (
	"github.com/tobycroft/Calc"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/FriendListModel"
	"main.go/app/bot/redis/FriendListRedis"
)

func App_refresh_friend_list_all() {
	bots := BotModel.Api_select()
	for _, bot := range bots {
		App_refresh_friend_list(Calc.Any2Int64(bot["self_id"]))
	}
}

func App_refresh_friend_list(self_id int64) {
	fl, err := iapi.Api.Getfriendlist(self_id)
	if err != nil {

	} else {
		App_refresh_friend_list_action(self_id, fl)
	}
}

func App_refresh_friend_list_action(self_id int64, fl []iapi.FriendList) {
	FriendListRedis.Cac_del(self_id, "*")
	FriendListModel.Api_delete(self_id)
	var fss []FriendListModel.FriendList
	for _, fll := range fl {
		var fs FriendListModel.FriendList
		fs.SelfId = self_id
		fs.UserId = fll.UserId
		fs.Nickname = fll.UserName
		fs.Remark = fll.UserRemark
		fss = append(fss, fs)
		FriendListRedis.Cac_set(self_id, fs.UserId, fs)
	}
	FriendListModel.Api_insert_more(fss)
}
