package BotDefaultReplyModel

import (
	"github.com/tobycroft/gorose"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "bot_default_reply"

func Api_select() []gorose.Data {
	db := tuuz.Db().Table(table)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
