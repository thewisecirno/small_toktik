package router

import (
	"SmallDouyin/config"
	"SmallDouyin/middleware"
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

		gr.POST("/publish/action/", middleware.JwtMiddleware(), service.PublishAction)
		gr.GET("/publish/list/", middleware.JwtMiddleware(), service.UserPublishList)

		gr.POST("/favorite/action/", middleware.JwtMiddleware(), service.UserFavoriteAction)
		gr.GET("/favorite/list/", middleware.JwtMiddleware(), service.UserFavoriteList)

		gr.POST("/comment/action/", middleware.JwtMiddleware(), service.UserCommentAction)
		gr.GET("/comment/list/", middleware.JwtMiddleware(), service.GetCommentList)

		gr.POST("/relation/action/", middleware.JwtMiddleware(), service.UserRelationAction)
		gr.GET("/relation/follow/list/", middleware.JwtMiddleware(), service.UserRelationFollowList)
		gr.GET("/relation/follower/list/", middleware.JwtMiddleware(), service.UserRelationFollowerList)

		//后面的东西前端都没有实现，那我也没办法了
		gr.GET("/relation/friend/list/", nil)

		gr.POST("/message/chat/", nil)
		gr.GET("/message/action/", nil)
	}
	return r
}
