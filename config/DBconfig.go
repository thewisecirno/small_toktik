package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func InitConfig() (err error) {
	viper.SetConfigFile("./config/config.yaml")
	err = viper.ReadInConfig()
	if err != nil {
		log.Println(err)
		return err
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Printf("change")
	})

	MysqlConf = MysqlConfig{
		User:     viper.GetString("mysql.user"),
		Password: viper.GetString("mysql.password"),
		IP:       viper.GetString("mysql.ip"),
		Port:     viper.GetString("mysql.port"),
		DbName:   viper.GetString("mysql.dbname"),
	}

	RedisConf = RedisConfig{
		Password: viper.GetString("redis.password"),
		IP:       viper.GetString("redis.ip"),
		Port:     viper.GetString("redis.port"),
		Db:       viper.GetString("redis.db"),
	}

	StaticSaveConf = StaticSavePath{
		Dst:     viper.GetString("save_path.dst"),
		DstName: viper.GetString("save_path.dst_name"),
	}

	LocalConf = Local{
		IP:   viper.GetString("local_config.wlan_ip"),
		Port: viper.GetString("local_config.port"),
	}
	return err
}

func InitMysql() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		MysqlConf.User,
		MysqlConf.Password,
		MysqlConf.IP,
		MysqlConf.Port,
		MysqlConf.DbName,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}

func InitRedis() (err error) {
	addr := fmt.Sprintf("%s:%s",
		RedisConf.IP,
		RedisConf.Port,
	)
	dbNum, err := strconv.Atoi(RedisConf.Db)
	if err != nil {
		log.Println("strconv Atoi dbNum error")
		return err
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: RedisConf.Password,
		DB:       dbNum,
	})
	return nil
}
