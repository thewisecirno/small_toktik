package router

import (
	"SmallDouyin/config"
	"SmallDouyin/service"
	"github.com/gin-gonic/gin"
	"log"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Static("static", config.StaticSaveConf.Dst)
	log.Println(config.StaticSaveConf.Dst)
	gr := r.Group("/douyin")
	{
		gr.GET("/feed", service.VideoFeed)

		gr.POST("/user/register/", service.UserRegister)
		gr.POST("/user/login/", service.UserLogin)
		gr.GET("/user/", service.GetUserInfo)

		gr.POST("/publish/action/", service.PublishAction)
		gr.GET("/publish/list/", service.UserPublishList)

		gr.POST("/favorite/action/", service.UserFavoriteAction)
		gr.GET("/favorite/list/", service.UserFavoriteList)

		gr.POST("/comment/action/", service.UserCommentAction)
		gr.GET("/comment/list/", service.GetCommentList)

		gr.POST("/relation/action/", service.UserRelationAction)
		gr.POST("/relation/follow/list/", service.UserRelationFollowList)
		gr.POST("/relation/follower/list/", nil)
		gr.POST("/relation/friend/list/", nil)

		gr.POST("/message/chat/", nil)
		gr.POST("/message/action/", nil)
	}
	return r
}
