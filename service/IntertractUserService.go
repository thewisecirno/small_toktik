package service

import (
	"SmallDouyin/middleware"
	"SmallDouyin/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

func UserFavoriteAction(c *gin.Context) {
	token := c.Query("token")
	videoIDStr := c.Query("video_id")
	actionType := c.Query("action_type")
	videoIdInt, err := strconv.ParseInt(videoIDStr, 10, 64)
	log.Println("video_id >>>>>>", videoIdInt)
	if err != nil {
		log.Println(err)
		return
	}

	parseToken, err := middleware.ParseToken(token)
	if err != nil {
		log.Println(err)
		return
	}

	if actionType == "1" {
		if models.IsUserFavorVideo(parseToken.UserInfoID, videoIdInt) {
			interactResponseFailed(c, "favor repeat")
			return
		}
		models.AscFavoriteCountById(videoIdInt)
		models.AddUserFavorVideos(parseToken.UserInfoID, videoIdInt)
		interactResponseSuccessfully(c, "favor successfully")
	} else if actionType == "2" {
		if !models.IsUserFavorVideo(parseToken.UserInfoID, videoIdInt) {
			interactResponseFailed(c, "favor quit repeat")
			return
		}
		models.DscFavoriteCountById(videoIdInt)
		models.RmUserFavorVideos(parseToken.UserInfoID, videoIdInt)
		interactResponseSuccessfully(c, "favor quit")
	} else {
		interactResponseFailed(c, "action_type undefined")
	}
}

func UserFavoriteList(c *gin.Context) {
	userIdStr := c.Query("user_id")
	userIdInt, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println(err)
		interactResponseFailed(c, "获取失败")
		return
	}
	videoList := models.GetUserFavorList(userIdInt)
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "获取成功",
		"video_list":  videoList,
	})
}

func UserCommentAction(c *gin.Context) {
	token := c.Query("token")
	videoIdStr := c.Query("video_id")
	actionType := c.Query("action_type")
	commentText := c.Query("comment_text")
	CommentIdStr := c.Query("comment_id")

	parseToken, err := middleware.ParseToken(token)
	if err != nil {
		log.Println(err)
		interactResponseFailed(c, "添加评论失败")
		return
	}

	videoIdInt, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		log.Println(err)
		interactResponseFailed(c, "添加评论失败")
		return
	}

	newComment := models.Comment{
		UserInfoId: parseToken.UserInfoID,
		VideoId:    videoIdInt,
		Content:    commentText,
		CreatedAt:  time.Time{},
		CreateDate: time.Now().String(),
	}

	if actionType == "1" {
		models.AddNewComment(newComment)
		models.AscCommentCountById(videoIdInt)
		c.JSON(http.StatusOK, gin.H{
			"status_code": 0,
			"status_msg":  "string",
			"comment":     newComment,
		})
	} else if actionType == "2" {
		CommentIdInt, err1 := strconv.ParseInt(CommentIdStr, 10, 64)
		if err1 != nil {
			log.Println(err1)
			interactResponseFailed(c, "删除评论失败")
			return
		}
		models.RemoveComment(CommentIdInt)
		models.DscCommentCountById(videoIdInt)
		interactResponseSuccessfully(c, "删除评论成功")
	} else {
		interactResponseFailed(c, "操作类型不存在")
	}
}

func GetCommentList(c *gin.Context) {
	videoIdStr := c.Query("video_id")
	videoIdInt, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		log.Println(err)
		interactResponseFailed(c, "获取评论列表失败")
		return
	}
	commentList := models.GetAllCommentByVid(videoIdInt)
	c.JSON(http.StatusOK, gin.H{
		"status_code":  0,
		"status_msg":   "获取评论列表成功",
		"comment_list": commentList,
	})
}

func interactResponseFailed(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  msg,
	})
}

func interactResponseSuccessfully(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  msg,
	})
}
