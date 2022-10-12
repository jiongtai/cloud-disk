package models

import (
	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

var Engine = Init()

var RDB = InitRedis()

func Init() *xorm.Engine {
	engine, err := xorm.NewEngine("mysql", "root:123456@tcp(127.0.0.1:3306)/cloud-disk?charset=utf8")
	if err != nil {
		log.Printf("Xorm New Engine Error:%v", err)
		return nil
	}
	return engine
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
}
