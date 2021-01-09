package GroupFunctionDetailModel

import (
	"main.go/tuuz"
	"main.go/tuuz/Log"
)

const table = "group_function_detail"

func Api_find_byK(k interface{}) interface{} {
	db := tuuz.Db().Table(table)
	where := map[string]interface{}{
		"k": k,
	}
	db.Where(where)
	ret, err := db.Value("v")
	if err != nil {
		Log.Dbrr(err, tuuz.FUNCTION_ALL())
		return nil
	} else {
		return ret
	}
}
