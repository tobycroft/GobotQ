package iapi

import (
	"errors"
	"github.com/bytedance/sonic"
	"github.com/gorilla/websocket"
	"github.com/tobycroft/Calc"
	Net "github.com/tobycroft/TuuzNet"
	"main.go/app/bot/model/BotModel"

	"main.go/tuuz/Log"
)

type UserInfoRet struct {
	Data    UserInfo `json:"data"`
	Retcode int64    `json:"retcode"`
	Status  string   `json:"status"`
}

type UserInfo struct {
	UserId    string      `json:"user_id"`
	Nickname  string      `json:"nickname"`
	Age       int64       `json:"age"`
	Sex       string      `json:"sex"`
	Level     int64       `json:"level"`
	LoginDays int64       `json:"login_days"`
	Qid       string      `json:"qid"`
	Vote      int64       `json:"vote"`
	WzryHonor interface{} `json:"wzry_honor"`
	Ext       struct {
		AddSrcId                   int64       `json:"add_src_id"`
		AddSrcName                 string      `json:"add_src_name"`
		AddSubSrcId                int64       `json:"add_sub_src_id"`
		AllowCalInteractive        bool        `json:"allow_cal_interactive"`
		AllowClick                 bool        `json:"allow_click"`
		AllowPeopleSee             bool        `json:"allow_people_see"`
		AuthState                  int64       `json:"auth_state"`
		BigClubVipOpen             int64       `json:"big_club_vip_open"`
		HollywoodVipOpen           int64       `json:"hollywood_vip_open"`
		QqVipOpen                  int64       `json:"qq_vip_open"`
		SuperQqOpen                int64       `json:"super_qq_open"`
		SuperVipOpen               int64       `json:"super_vip_open"`
		Voted                      int64       `json:"voted"`
		BabyQSwitch                bool        `json:"baby_q_switch"`
		BindPhoneInfo              string      `json:"bind_phone_info"`
		CardId                     int64       `json:"card_id"`
		CardType                   int64       `json:"card_type"`
		Category                   int64       `json:"category"`
		ClothesId                  int64       `json:"clothes_id"`
		CoverUrl                   string      `json:"cover_url"`
		Declaration                interface{} `json:"declaration"`
		DefaultCardId              int64       `json:"default_card_id"`
		DiyComplicatedInfo         interface{} `json:"diy_complicated_info"`
		DiyDefaultText             interface{} `json:"diy_default_text"`
		DiyText                    interface{} `json:"diy_text"`
		DiyTextDegree              float64     `json:"diy_text_degree"`
		DiyTextFontId              int64       `json:"diy_text_font_id"`
		DiyTextHeight              float64     `json:"diy_text_height"`
		DiyTextWidth               float64     `json:"diy_text_width"`
		DiyTextLocX                float64     `json:"diy_text_loc_x"`
		DiyTextLocY                float64     `json:"diy_text_loc_y"`
		DressUpIsOn                bool        `json:"dress_up_is_on"`
		EncId                      interface{} `json:"enc_id"`
		EnlargeQzonePic            int64       `json:"enlarge_qzone_pic"`
		ExtendFriendEntryAddFriend int64       `json:"extend_friend_entry_add_friend"`
		ExtendFriendEntryContact   int64       `json:"extend_friend_entry_contact"`
		ExtendFriendFlag           int64       `json:"extend_friend_flag"`
		ExtendFriendQuestion       int64       `json:"extend_friend_question"`
		ExtendFriendVoiceDuration  int64       `json:"extend_friend_voice_duration"`
		FavoriteSource             int64       `json:"favorite_source"`
		FeedPreviewTime            int64       `json:"feed_preview_time"`
		FontId                     int64       `json:"font_id"`
		FontType                   int64       `json:"font_type"`
		QidBgUrl                   string      `json:"qid_bg_url"`
		QidColor                   string      `json:"qid_color"`
		QidLogoUrl                 string      `json:"qid_logo_url"`
		QqCardIsOn                 bool        `json:"qq_card_is_on"`
		SchoolId                   interface{} `json:"school_id"`
		SchoolName                 interface{} `json:"school_name"`
		SchoolVerifiedFlag         bool        `json:"school_verified_flag"`
		ShowPublishButton          bool        `json:"show_publish_button"`
		Singer                     string      `json:"singer"`
		SongDura                   int64       `json:"song_dura"`
		SongId                     string      `json:"song_id"`
		SongName                   string      `json:"song_name"`
	} `json:"ext"`
}

func (api Post) GetStrangerInfo(self_id, user_id int64, no_cache bool) (UserInfo, error) {
	post := map[string]any{
		"user_id":  user_id,
		"no_cache": no_cache,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return UserInfo{}, errors.New("botinfo_notfound")
	}
	data, err := Net.Post{}.PostUrlXEncode(botinfo["url"].(string)+"/get_stranger_info", nil, post, nil, nil).RetString()
	if err != nil {
		return UserInfo{}, err
	}
	var ret1 UserInfoRet

	err = sonic.UnmarshalString(data, &ret1)
	if err != nil {
		return UserInfo{}, err
	}
	return ret1.Data, nil
}
func (api Ws) GetStrangerInfo(self_id, user_id int64, no_cache bool) (UserInfo, error) {
	post := map[string]any{
		"user_id":  user_id,
		"no_cache": no_cache,
	}
	botinfo := BotModel.Api_find(self_id)
	if len(botinfo) < 1 {
		Log.Crrs(nil, "bot:"+Calc.Any2String(self_id))
		return UserInfo{}, errors.New("botinfo_notfound")
	}
	data, err := sonic.Marshal(sendStruct{
		Action: "get_stranger_info",
		Params: post,
		Echo: echo{
			Action: "get_stranger_info",
			SelfId: Calc.Any2Int64(self_id),
		},
	})
	if err != nil {
		return UserInfo{}, err
	}
	conn, ok := ClientToConn.Load(self_id)
	if !ok {
		return UserInfo{}, errors.New("ClientNotFound")
	}
	Net.WsServer_WriteChannel <- Net.WsData{
		Conn: conn.(*websocket.Conn), Message: data,
	}
	return UserInfo{}, err
}
