package service

import (
	"SmallDouyin/middleware"
	"SmallDouyin/models"
	"SmallDouyin/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func GetUserInfo(c *gin.Context) {
	userIdStr := c.Query("user_id")
	log.Println(userIdStr)
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		log.Println("userId解析失败")
		responseUserBasicFailed(c, "userId解析失败")
		return
	}
	userInfo := models.FindUserInfoById(userId)
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "获取用户信息成功",
		"user":        userInfo,
	})
}

func UserRegister(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	fmt.Println(">>>", username, password)
	userLogin := models.FindUsername(username)
	if userLogin.Username != "" {
		responseUserBasicFailed(c, "register failed,repeat")
		return
	} else {
		userLogin.Username = username
		userLogin.Password = utils.EncryptPassword(password)
		userID := models.CreateUserInfo(userLogin)
		userLogin.UserInfoId = userID
		models.CreateUserLogin(userLogin)
		token, err := middleware.GetToken(userLogin.Username, userLogin.Password, userID)
		if err != nil {
			responseUserBasicFailed(c, "register failed,can't get token")
			return
		}
		responseUserBasicSuccessful(c, userID, token, "register successfully")
	}
}

func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	userLogin := models.FindUsername(username)
	if userLogin.Username == "" {
		log.Println("用户名或者密码错误")
		responseUserBasicFailed(c, "用户名或者密码错误")
		return
	} else {
		if !utils.DecryptPassword(password, userLogin.Password) {
			log.Println("用户名或者密码错误")
			responseUserBasicFailed(c, "用户名或者密码错误")
			return
		}
		token, err := middleware.GetToken(username, password, userLogin.UserInfoId)

		if err != nil {
			log.Println(err)
			responseUserBasicFailed(c, "登陆成功，但是token生成失败")
			return
		}
		responseUserBasicSuccessful(c, userLogin.Id, token, "login successfully")
	}
}

func responseUserBasicSuccessful(c *gin.Context, userID int64, token, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  msg,
		"user_id":     userID,
		"token":       token,
	})
}

func responseUserBasicFailed(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"status_code": -1,
		"status_msg":  msg,
		"user_id":     -1,
		"token":       "",
	})
}
