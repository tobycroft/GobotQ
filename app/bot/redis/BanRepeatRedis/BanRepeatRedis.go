package BanRepeatRedis

import (
	"github.com/tobycroft/Calc"
	"main.go/tuuz/Redis"
	"strings"
	"time"
)

const FriendList = "BanRepeat"
const UserId = ":UserId:"
const Message = ":Message:"

type BanRepeatRedis struct {
	table string
}

func (self BanRepeatRedis) Table(user_id, message any) BanRepeatRedis {
	str := strings.Builder{}
	str.WriteString(FriendList)
	str.WriteString(UserId)
	if user_id != nil {
		str.WriteString(Calc.Any2String(user_id))
	} else {
		str.WriteString("*")
	}
	str.WriteString(Message)
	if message != nil {
		str.WriteString(Calc.Md5(Calc.Any2String(message)))
	} else {
		str.WriteString("*")
	}
	self.table = str.String()
	return self
}

func (self BanRepeatRedis) Cac_set(data int64, exp time.Duration) error {
	return Redis.String_set(self.table, data, exp)
}

func (self BanRepeatRedis) Cac_find() (int64, error) {
	return Redis.String_getInt64(self.table)
}
