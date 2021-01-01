package api

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"main.go/config/app_conf"
	"main.go/tuuz/Net"
)

type Gls struct {
	Ret  string    `json:"ret"`
	List GroupList `json:"List"`
}

type GroupList []struct {
	GroupID                int    `json:"GroupID"`
	GIN                    int    `json:"GIN"`
	CFlag                  int    `json:"cFlag"`
	GroupInfoSeq           int    `json:"GroupInfoSeq"`
	DwGroupFlagExt         int    `json:"dwGroupFlagExt"`
	DwGroupRankSeq         int    `json:"dwGroupRankSeq"`
	DwCertificationType    int    `json:"dwCertificationType"`
	DwShutupTimestamp      int    `json:"dwShutupTimestamp"`
	DwMyShutupTimestamp    int    `json:"dwMyShutupTimestamp"`
	DwCmdUinUinFlag        int    `json:"dwCmdUinUinFlag"`
	DwAdditionalFlag       int    `json:"dwAdditionalFlag"`
	DwGroupTypeFlag        int    `json:"dwGroupTypeFlag"`
	DwGroupSecType         int    `json:"dwGroupSecType"`
	DwGroupSecTypeInfo     int    `json:"dwGroupSecTypeInfo"`
	DwGroupClassExt        int    `json:"dwGroupClassExt"`
	DwAppPrivilegeFlag     int    `json:"dwAppPrivilegeFlag"`
	DwSubscriptionUin      int    `json:"dwSubscriptionUin"`
	DwMemberNum            int    `json:"dwMemberNum"`
	DwMemberNumSeq         int    `json:"dwMemberNumSeq"`
	DwMemberCardSeq        int    `json:"dwMemberCardSeq"`
	DwGroupFlagExt3        int    `json:"dwGroupFlagExt3"`
	DwGroupOwnerUin        int    `json:"dwGroupOwnerUin"`
	CIsConfGroup           int    `json:"cIsConfGroup"`
	CIsModifyConfGroupFace int    `json:"cIsModifyConfGroupFace"`
	CIsModifyConfGroupName int    `json:"cIsModifyConfGroupName"`
	DwCmduinJoinTime       int    `json:"dwCmduinJoinTime"`
	StrGroupName           string `json:"strGroupName"`
	StrGroupMemo           string `json:"strGroupMemo"`
}

func Getgrouplist(bot interface{}) (GroupList, error) {
	post := map[string]interface{}{
		"logonqq": bot,
	}
	data, err := Net.Post(app_conf.Http_Api+"/getgrouplist", nil, post, nil, nil)
	fmt.Println(data)
	if err != nil {
		return nil, err
	}
	var gls Gls
	jsr := jsoniter.ConfigCompatibleWithStandardLibrary
	err = jsr.UnmarshalFromString(data, &gls)
	if err != nil {
		return nil, err
	}
	return gls.List, nil
}
