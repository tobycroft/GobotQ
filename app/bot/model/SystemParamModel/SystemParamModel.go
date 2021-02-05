package SystemParamModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "system_param"

func Api_value(key interface{}) interface{} {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"key": key,
	}
	db.Where(where)
	ret, err := db.Value("value")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
