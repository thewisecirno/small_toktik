package models

import (
	"SmallDouyin/config"
	"gorm.io/gorm"
	"strconv"
)

type UserFavorVideos struct {
	UserInfoId int64
	VideoId    int64
}

func AddUserFavorVideos(userInfoId, VideoId int64) {
	userFavorVideos := UserFavorVideos{
		UserInfoId: userInfoId,
		VideoId:    VideoId,
	}
	config.DB.Create(&userFavorVideos)
}

func RmUserFavorVideos(userInfoId, videoId int64) {
	config.DB.Where("user_info_id = ? and video_id = ?", userInfoId, videoId).Delete(&UserFavorVideos{})
}

func IsUserFavorVideo(userInfoId, VideoId int64) bool {
	if config.DB.
		Where("user_info_id = ? and video_id = ?", userInfoId, VideoId).
		First(&UserFavorVideos{}).Error == gorm.ErrRecordNotFound {
		return false
	}
	return true
}

func GetUserFavorVideoId(userId int64) []string {
	var userFavorVideosList []*UserFavorVideos
	var favorVideoIdList []string
	config.DB.Where("user_info_id = ? ", userId).Find(&userFavorVideosList)
	for _, v := range userFavorVideosList {
		videoId := strconv.FormatInt(v.VideoId, 10)
		favorVideoIdList = append(favorVideoIdList, videoId)
	}
	return favorVideoIdList
}

func GetUserFavorList(userID int64) []*Video {
	var FavorVideoList []*Video
	VideoIdList := GetUserFavorVideoId(userID)
	config.DB.Where("id in ?", VideoIdList).Find(&FavorVideoList)
	return FavorVideoList
}
