package iapi

type IfaceApi interface {
	DeleteFriend(self_id, friend_id any) (bool, error)
	DeleteMsg(self_id, message_id any) (bool, error)
	Getfriendlist(self_id any) ([]FriendList, error)
	GetGroupInfo(self_id, group_id any) (GroupInfo, error)
	Getgrouplist(self_id any) ([]GroupList, error)
	GetGroupMemberInfo(self_id, group_id, user_id any) (GroupMemberList, error)
	Getgroupmemberlist(self_id, group_id any) ([]GroupMemberList, error)
	GetLoginInfo(self_id any) (LoginInfo, error)
	GetStrangerInfo(self_id, user_id any, no_cache bool) (UserInfo, error)
	Sendgroupmsg(Self_id, Group_id any, Message string, AutoRetract bool)
	Sendprivatemsg(Self_id, UserId, GroupId any, Message string, AutoRetract bool)
	SetFriendAddRequest(self_id, flag any, approve bool, remark any) (bool, error)
	SetGroupAddRequestRet(self_id, flag, sub_type any, approve bool, reason string) (bool, error)
	SetGroupAdmin(self_id, group_id, user_id any, enable bool) (bool, error)
	SetGroupBan(self_id, group_id, user_id any, duration float64) (bool, error)
	Setgroupcard(self_id, group_id, user_id, card any) (bool, error)
	SetGroupKick(self_id, group_id, user_id any, reject_add_request bool) (bool, error)
	SetGroupLeave(self_id, group_id any) (bool, error)
	SetGroupWholeBan(self_id, group_id any, enable bool) (bool, error)
}
