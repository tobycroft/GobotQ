package GroupMemberModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_member"

type GroupMember struct {
	bot        interface{} `gorose:"uid"`
	gid        int         `gorose:"gid"`
	uid        int         `gorose:"uid"`
	age        int         `gorose:"age"`
	title      string      `gorose:"title"`
	nickname   string      `gorose:"nickname"`
	remark     string      `gorose:"remark"`
	card       string      `gorose:"card"`
	jointime   int         `gorose:"jointime"`
	lastsend   int         `gorose:"lastsend"`
	grouplevel int         `gorose:"grouplevel"`
}

func Api_insert(gm GroupMember) bool {
	db := tuuz.Db().Table(table)
	db.Data(gm)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_insert_more(gms []GroupMember) bool {
	db := tuuz.Db().Table(table)
	db.Data(gms)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_select(bot) {

}
