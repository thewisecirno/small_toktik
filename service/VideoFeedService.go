package service

import (
	"SmallDouyin/middleware"
	"SmallDouyin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func VideoFeed(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		DoWithoutToken(c)
	} else {
		DoWithToken(c, token)
	}
}

func DoWithToken(c *gin.Context, token string) {
	//客制化内容，随便写点东西吧
	//这里就随便写个，无法刷到自己投递的内容
	TimeStampStr := c.Query("latest_time")
	TimeStampInt, err := strconv.ParseInt(TimeStampStr, 10, 64)
	log.Println("TimeStampInt >>>>>", TimeStampInt)
	if err != nil {
		log.Println(err)
		return
	}

	parseToken, err := middleware.ParseToken(token)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("%+v", parseToken)
	videoList := models.GetVideoListWithoutSelf(time.Unix(0, TimeStampInt*1e6), parseToken.UserInfoID)
	fmt.Printf(" videoList >>>>> %+v\n", videoList)

	var nextTime interface{}
	if len(videoList) == 0 {
		nextTime = nil
	} else {
		nextTime = videoList[0].CreatedAt.Unix()
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "get user's publish-list successfully",
		"next_time":   nextTime,
		"video_list":  videoList,
	})
}

func DoWithoutToken(c *gin.Context) {
	TimeStampStr := c.Query("latest_time")
	TimeStampInt, err := strconv.ParseInt(TimeStampStr, 10, 64)
	log.Println("TimeStampInt >>>>>", TimeStampInt)
	if err != nil {
		log.Println(err)
		return
	}

	videoList := models.GetVideoList(time.Unix(0, TimeStampInt*1e6))
	fmt.Printf(" videoList >>>>> %+v\n", videoList)

	var nextTime interface{}
	if len(videoList) == 0 {
		nextTime = nil
	} else {
		nextTime = videoList[0].CreatedAt.Unix()
	}
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "get user's publish-list successfully",
		"next_time":   nextTime,
		"video_list":  videoList,
	})
}
