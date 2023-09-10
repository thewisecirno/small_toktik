package service

import (
	"SmallDouyin/middleware"
	"SmallDouyin/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func UserRelationAction(c *gin.Context) {
	token := c.Query("token")
	toUserIdStr := c.Query("to_user_id")
	actionType := c.Query("action_type")
	toUserIdInt, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		log.Println(err)
		responseUserRelationFailed(c, "error")
		return
	}

	parseToken, err := middleware.ParseToken(token)
	if err != nil {
		log.Println(err)
		responseUserRelationFailed(c, "token error")
		return
	}
	if actionType == "1" {
		models.AddRelation(parseToken.UserInfoID, toUserIdInt)
		responseUserRelationSuccessfully(c, "关注成功")

	} else if actionType == "2" {
		models.RemoveRelation(parseToken.UserInfoID, toUserIdInt)
		responseUserRelationSuccessfully(c, "取消关注成功")

	} else {
		responseUserRelationFailed(c, "未知操作")
	}
}

func UserRelationFollowList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	//token := c.Query("token")
	userIdInt, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println(err)
		responseUserRelationFailed(c, "parse int64 error")
		return
	}
	UserFollowList := models.GetUserFollowList(userIdInt)
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "get user follow list successfully",
		"user_list":   UserFollowList,
	})
}

func responseUserRelationFailed(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": -1,
		"status_msg":  msg,
	})
}

func responseUserRelationSuccessfully(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  msg,
	})
}
