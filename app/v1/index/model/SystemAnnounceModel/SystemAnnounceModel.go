package SystemAnnounceModel

import (
	"github.com/tobycroft/gorose-pro"
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "system_announce"

func Api_select() []gorose.Data {
	db := tuuz.Db().Table(table)
	where := map[string]any{
		"type": 1,
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
