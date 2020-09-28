package initialize

import (
	"github.com/HuLuWang/blog_demo/global"
	"github.com/HuLuWang/blog_demo/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
)

func Logger() {
	fileName := global.CONFIG.Logger.LogSavePath + "/" + global.CONFIG.Logger.LogFileName + global.CONFIG.Logger.LogFileExt
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "", log.LstdFlags).WithCaller(2)
}
