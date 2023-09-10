package service

import (
	"SmallDouyin/config"
	"SmallDouyin/middleware"
	"SmallDouyin/models"
	"SmallDouyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

var (
	VideoFormat = map[string]interface{}{
		".mp4": nil,
		".avi": nil,
		".mov": nil,
		".wmv": nil,
		".flv": nil,
	}
)

func PublishAction(c *gin.Context) {
	title := c.PostForm("title")
	token := c.PostForm("token")
	formData, err := c.MultipartForm()
	if err != nil {
		log.Println(err)
		responseVideoFailed(c, "publish failed")
		return
	}
	file := formData.File["data"]
	fmt.Println(title, token)
	fmt.Printf(">>>>>%#v\n", formData.File)
	fmt.Printf(">>>>>%#v\n", formData.Value)
	for _, v := range file {
		suffix := filepath.Ext(v.Filename)
		if _, ok := VideoFormat[suffix]; !ok {
			log.Println("不支持的视频格式")
			responseVideoFailed(c, "不支持的视频格式")
			return
		}
		var parseToken *middleware.UserBasicClaims
		parseToken, err = middleware.ParseToken(token)
		if err != nil {
			log.Println(err)
			responseVideoFailed(c, "publish failed")
			return
		}
		filename := utils.GetVideoFileName(parseToken.UserInfoID) + suffix
		savePath := filepath.Join(config.StaticSaveConf.Dst, filename)
		fmt.Println(">>>>>", savePath)
		err = c.SaveUploadedFile(v, savePath)

		PublishVideo := models.Video{
			UserInfoId: parseToken.UserInfoID,
			Author:     models.FindUserInfoById(parseToken.UserInfoID),
			PlayUrl:    utils.GetVideoURL(filename),
			CoverUrl:   "",
			Title:      title,
		}
		models.CreateVideo(PublishVideo)

		if err != nil {
			log.Println("save file error", err)
			responseVideoFailed(c, "publish failed")
			return
		}

		responseVideoSuccessful(c, "publish successfully")
	}
}

func UserPublishList(c *gin.Context) {
	userId, err := strconv.ParseInt(c.Query("user_id"), 10, 64)
	if err != nil {
		log.Println(err)
		responseVideoFailed(c, "get user's publish-list failed")
		return
	}
	videoList := models.FindVideoByUid(userId)

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "get user's publish-list successfully",
		"video_list":  videoList,
	})
}

func responseVideoSuccessful(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  msg,
	})
}

func responseVideoFailed(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": -1,
		"status_msg":  msg,
	})
}
