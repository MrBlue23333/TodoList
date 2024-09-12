package main

import (
	"demo/config"
	"demo/pkg/utils"
	"demo/repository/cache"
	"demo/repository/db/dao"
	"demo/router"
)

func main() {
	loadDb()
	r := router.NewRouter()
	_ = r.Run(config.Config.System.HttpPort)
}

func loadDb() {
	config.InitConfig()
	dao.MySQLInit()
	utils.InitLog()
	cache.RedisInit()
}
