package GroupMemberModel

import (
	"github.com/gohouse/gorose/v2"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_member"

type GroupMember struct {
	Bot        interface{} `gorose:"uid"`
	Gid        interface{} `gorose:"gid"`
	Uid        interface{} `gorose:"uid"`
	Title      string      `gorose:"title"`
	Nickname   string      `gorose:"nickname"`
	Remark     string      `gorose:"remark"`
	Card       string      `gorose:"card"`
	Jointime   int         `gorose:"jointime"`
	Lastsend   int         `gorose:"lastsend"`
	Grouplevel int         `gorose:"grouplevel"`
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

func Api_select(bot, gid interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot": bot,
		"gid": gid,
	}
	db.Where(where)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_delete_byGid(bot, gid interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot": bot,
		"gid": gid,
	}
	db.Where(where)
	_, err := db.Delete()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_delete(bot interface{}) bool {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot": bot,
	}
	db.Where(where)
	_, err := db.Delete()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}

func Api_find(bot, gid, uid interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot": bot,
		"gid": gid,
		"uid": uid,
	}
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_find_owner(bot, gid interface{}) gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot":        bot,
		"gid":        gid,
		"grouplevel": 6,
	}
	db.Where(where)
	ret, err := db.First()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}

func Api_select_admin(bot, gid interface{}) []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"bot":        bot,
		"gid":        gid,
		"grouplevel": 4,
	}
	db.Where(where)
	ret, err := db.Get()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
