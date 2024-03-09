package group

import (
	"fmt"
	"github.com/bytedance/sonic"
	"main.go/app/bot/model/GroupFunctionModel"
	"main.go/app/bot/model/GroupMemberModel"
	"main.go/app/bot/service"
	"main.go/config/types"
	"main.go/tuuz/Log"
	"main.go/tuuz/Redis"
)

func group_message_acfur_semi_match() {
	go ban_group()
	go ban_url()
	go ban_wx()
	go ban_share()

	go ban_word()
	go set_setting()
	go sign_daily()
	go sign_lunpan()

	go check_score()
	go rank_list()
	go word_limit()

	go re_verify()
	go ad_verify()

	go auto_reply()
	go greeting_when_at_me()
	go daoju()
	go trade()
	go pal()

	ps := Redis.PubSub{}
	for c := range ps.Subscribe(types.MessageGroupNormal) {
		fmt.Println("semi")
		var es EventStruct[GroupMessageStruct]
		err := sonic.UnmarshalString(c.Payload, &es)
		if err != nil {
			Log.Err(err)
		} else {
			gm := es.Json

			//self_id := gm.SelfId
			user_id := gm.UserId
			group_id := gm.GroupId
			//message_id := gm.MessageId
			//message := gm.Message
			message := gm.RawMessage
			raw_message := gm.RawMessage

			//text := message
			//reg := regexp.MustCompile("(?i)^acfur")
			//active := reg.MatchString(text)
			//new_text := reg.ReplaceAllString(text, "")

			groupmember := GroupMemberModel.Api_find(group_id, user_id)
			groupfunction := GroupFunctionModel.Api_find(group_id)
			if len(groupfunction) < 1 {
				GroupFunctionModel.Api_insert(group_id)
				groupfunction = GroupFunctionModel.Api_find(group_id)
			}
			gmr := GroupMessageRedirect[GroupMessageStruct]{}
			gmr.GroupMember = groupmember
			gmr.GroupFunction = groupfunction
			gmr.Json = gm

			if groupfunction["ban_group"].(int64) == 1 {
				if service.Serv_ban_group(raw_message) {
					//fmt.Println(banGroup, self_id, group_id, user_id)
					ps.Publish_struct(types.MessageGroupAcfur+banGroup, gmr)
					continue
				}
			}

			if groupfunction["ban_url"].(int64) == 1 {
				if service.Serv_url_detect(raw_message) {
					//fmt.Println(banUrl, self_id, group_id, user_id)
					ps.Publish_struct(types.MessageGroupAcfur+banUrl, gmr)
					continue
				}
			}

			if groupfunction["ban_wx"].(int64) == 1 {
				if service.Serv_ban_weixin(message) {
					//fmt.Println(banWx, self_id, group_id, user_id)
					ps.Publish(types.MessageGroupAcfur+banWx, gmr)
					continue
				}
			}

			if groupfunction["ban_share"].(int64) == 1 {
				if service.Serv_ban_share(message) {
					//fmt.Println(banShare, self_id, group_id, user_id)
					ps.Publish(types.MessageGroupAcfur+banShare, gmr)
					continue
				}
			}

			if groupfunction["sign"].(int64) == 1 {
				if _, ok := service.Serv_text_match_all(message, []string{"签到"}); ok {
					ps.Publish(types.MessageGroupAcfur+signDaily, gmr)
					continue
				}
				if _, ok := service.Serv_text_match(message, []string{"轮盘"}); ok {
					ps.Publish(types.MessageGroupAcfur+signLunpan, gmr)
					continue
				}
			}

			if _, ok := service.Serv_text_match_all(message, []string{"积分查询", "查询积分", "威望查询", "查询威望", "钱包", "查询余额", "余额查询"}); ok {
				ps.Publish(types.MessageGroupAcfur+checkScore, gmr)
			}

			if _, ok := service.Serv_text_match_all(message, []string{"积分排行", "威望排行", "排行榜"}); ok {
				ps.Publish(types.MessageGroupAcfur+rankList, gmr)
			}
		}
	}
}
