package LogsModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "logs"

func Api_insert(log, discript interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"log":      log,
		"discript": discript,
	}
	db.Data(data)
	_, err := db.Insert()
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return false
	} else {
		return true
	}
}
