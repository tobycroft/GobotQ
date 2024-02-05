package private

import (
	"errors"
	"github.com/tobycroft/Calc"
	"main.go/app/bot/iapi"
	"main.go/app/bot/model/BotModel"
	"main.go/app/bot/model/PrivateAutoReplyModel"
	"main.go/config/app_default"
	"main.go/tuuz/Log"
	"regexp"
	"time"
)

var PrivateMsgChan = make(chan PrivateMessageStruct, 99)

func (pm PrivateMessageStruct) PrivateHandle(selfId, user_id, group_id int64, message, rawMessage string) {
	reg := regexp.MustCompile("(?i)^acfur")
	active := reg.MatchString(message)
	new_text := reg.ReplaceAllString(message, "")

	botinfo := BotModel.Api_find(selfId)
	//if botinfo["allow_ip"] == nil {
	//	return
	//}
	//if !strings.Contains(remoteip, botinfo["allow_ip"].(string)) {
	//	Log.Errs(errors.New(fmt.Sprint(remoteip, botinfo["allow_ip"].(string))), "不允许的ip")
	//	return
	//}

	if len(botinfo) < 1 {
		Log.Crrs(errors.New("bot_not_found"), Calc.Any2String(selfId))
		return
	}
	if botinfo["end_date"].(time.Time).Before(time.Now()) {
		iapi.Api.Sendprivatemsg(selfId, user_id, group_id, app_default.Default_over_time, false)
		return
	}
	if active {
		pm.active_main_function(selfId, user_id, group_id, new_text, message)
	} else {
		//在未激活acfur的情况下应该对原始内容进行还原
		if private_default_reply(selfId, user_id, group_id, message) {
			return
		}
		auto_reply := PrivateAutoReplyModel.Api_find_byKey(message)
		if len(auto_reply) > 0 {
			if auto_reply["value"] == nil {
				return
			}
			iapi.Api.Sendprivatemsg(selfId, user_id, group_id, auto_reply["value"].(string), false)
		} else {
			private_auto_reply(selfId, user_id, group_id, message)
		}
	}
}
