package models

import (
	"SmallDouyin/config"
)

type UserLogin struct {
	Id         int64 `gorm:"primary_key"`
	UserInfoId int64
	Username   string `gorm:"primary_key"`
	Password   string `gorm:"password"`
}

func (user *UserLogin) TableName() string {
	return "user_login"
}

func FindUsername(username string) UserLogin {
	userLogin := UserLogin{}
	config.DB.Where("username = ?", username).First(&userLogin)
	return userLogin
}

func CreateUserLogin(userLogin UserLogin) int64 {
	config.DB.Create(&userLogin)
	return userLogin.Id
}
