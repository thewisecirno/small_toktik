package models

import (
	"SmallDouyin/config"
	"log"
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
	config.DB.Model(&UserRelations{}).Where("follow_id = ?", userId).Find(&userRelations)
	log.Println("userRelations", userRelations)
	for _, v := range userRelations {
		followId := strconv.FormatInt(v.UserInfoId, 10)
		userFollowId = append(userFollowId, followId)
	}
	log.Println("user follow id", userFollowId)
	return userFollowId
}

func GetUserFollowList(userId int64) []*UserInfo {
	var userFollowList []*UserInfo
	userFollowIdList := GetUserFollowId(userId)
	config.DB.Where("id in ?", userFollowIdList).Find(&userFollowList)
	return userFollowList
}

func GetUserFollowerId(userId int64) []string {
	var userFollowId []string
	var userRelations []*UserRelations
	config.DB.Model(&UserRelations{}).Where("user_info_id = ?", userId).Find(&userRelations)
	log.Println("userRelations", userRelations)
	for _, v := range userRelations {
		followId := strconv.FormatInt(v.UserInfoId, 10)
		userFollowId = append(userFollowId, followId)
	}
	log.Println("user follower id", userFollowId)
	return userFollowId
}

func GetUserFollowerList(userId int64) []*UserInfo {
	var userFollowList []*UserInfo
	userFollowIdList := GetUserFollowerId(userId)
	config.DB.Where("id in ?", userFollowIdList).Find(&userFollowList)
	return userFollowList
}
