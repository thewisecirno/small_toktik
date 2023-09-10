package main

import (
	"SmallDouyin/config"
	"SmallDouyin/models"
	"SmallDouyin/router"
	"log"
)

func main() {
	err := InitAll()
	if err != nil {
		log.Fatal(err)
	}

	r := router.InitRouter()
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}

func InitAll() (err error) {
	err = config.InitConfig()
	if err != nil {
		return err
	}

	err = config.InitMysql()
	if err != nil {
		return err
	}

	err = config.InitRedis()
	if err != nil {
		return err
	}

	err = config.DB.AutoMigrate(
		&models.UserInfo{},
		&models.Comment{},
		&models.Message{},
		&models.Video{},
		&models.UserLogin{})

	if err != nil {
		return err
	}
	return nil
}
