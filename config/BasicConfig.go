package config

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	Rdb *redis.Client
)

var (
	MysqlConf      MysqlConfig
	RedisConf      RedisConfig
	StaticSaveConf StaticSavePath
	LocalConf      Local
)

type MysqlConfig struct {
	User     string
	Password string
	IP       string
	Port     string
	DbName   string
}

type RedisConfig struct {
	Password string
	IP       string
	Port     string
	Db       string
}

type StaticSavePath struct {
	Dst     string
	DstName string
}

type Local struct {
	IP   string
	Port string
}
