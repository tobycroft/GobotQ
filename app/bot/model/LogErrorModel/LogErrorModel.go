package LogErrorModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "log_error"

func Api_insert(error, discript interface{}) bool {
	db := tuuz.Db().Table(table)
	data := map[string]interface{}{
		"error":    error,
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
