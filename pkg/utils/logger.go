package utils

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"time"
)

var LogrusObj *logrus.Logger

func InitLog() {
	if LogrusObj != nil {
		src, _ := setOutPutFile()
		LogrusObj.Out = src
		return
	}

	logger := logrus.New()
	src, _ := setOutPutFile()
	logger.Out = src
	logrus.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	LogrusObj = logger
}

func setOutPutFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
		fmt.Println(err)
	}
	//分别检查 日志文件夹和日志文件 路径是否存在，不存在则创建
	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, os.ModePerm); err != nil {
			return nil, err
		}
	}
	logFileName := now.Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	_, err = os.Stat(fileName)
	if os.IsNotExist(err) {
		if _, err := os.Create(fileName); err != nil {
			return nil, err
		}
	}
	//打开路径下的日志文件,返回文件描述符
	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend|os.ModePerm)
	if err != nil {
		return nil, err
	}
	return src, nil
}
