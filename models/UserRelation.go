package models

import (
	"SmallDouyin/config"
	"strconv"
)

type UserRelations struct {
	UserInfoId int64
	FollowId   int64
}

func AddRelation(userInfoId, followId int64) {
	userRelation := UserRelations{
		UserInfoId: userInfoId,
		FollowId:   followId,
	}
	config.DB.Create(&userRelation)
}

func RemoveRelation(userInfoId, followId int64) {
	config.DB.
		Where("user_info_id = ? and follow_id = ?", userInfoId, followId).
		Delete(&UserRelations{})
}

func GetUserFollowId(userId int64) []string {
	var userFollowId []string
	var userRelations []*UserRelations
	config.DB.Where("follow_id = ?", userId).Find(&userRelations)
	for _, v := range userRelations {
		followId := strconv.FormatInt(v.UserInfoId, 10)
		userFollowId = append(userFollowId, followId)
	}
	return userFollowId
}

func GetUserFollowList(userId int64) []*UserInfo {
	var userFollowList []*UserInfo
	userFollowIdList := GetUserFollowId(userId)
	config.DB.Where("user_info_id in ?", userFollowIdList).Find(&userFollowList)
	return userFollowList
}
