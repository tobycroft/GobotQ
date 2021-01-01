package api

import (
	"fmt"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

func Getgrouplist(bot interface{}) {
	post := map[string]interface{}{
		"logonqq": bot,
	}
	data, err := Net.Post(app_conf.Http_Api+"/getgrouplist", nil, post, nil, nil)
	fmt.Println(data)
	if err != nil {
		//return nil, err
	}
	//var gfl GFL
	//jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	//err = jsr.UnmarshalFromString(data, &gfl)
	//if err != nil {
	//	return nil, err
	//}
	//return gfl.List, nil
}
