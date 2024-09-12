package dao

import (
	"context"
	"demo/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

var _db *gorm.DB

func MySQLInit() {
	//选取config中的default mysql配置
	mConfig := config.Config.MySql["default"]
	conn := strings.Join([]string{mConfig.UserName, ":", mConfig.Password,
		"@tcp(", mConfig.DbHost, ":", mConfig.DbPort, ")/", mConfig.DbName,
		"?charset=", mConfig.Charset,
		"&parseTime=True&loc=Local"}, "")

	var ormLogger = logger.Default
	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       conn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger, //打印日志
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //不加s
		},
	})
	if err != nil {
		panic(err)
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db
	migrate()
}

// 返回上下文相关的GORM.DB
func NewDbClient(c context.Context) *gorm.DB {
	db := _db
	return db.WithContext(c)
}
