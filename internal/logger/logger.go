package logger

import (
	"Fank/internal/global"
	"github.com/sirupsen/logrus"
	"os"
)

func init() {

	initLogger()
}

//func initLogger() {
//	config, err := configs.LoadConfig()
//	if err != nil {
//		log.Fatalf("初始化日志组件时加载配置失败: %v", err)
//	}
//
//	logFilePath := config.LogConfig.LogFilePath
//	logFileName := config.LogConfig.LogFileName
//
//}

func initLogger() {
	// 初始化全局日志实例
	global.SysLog = logrus.New()
	global.SysLog.SetFormatter(&logrus.JSONFormatter{})
	global.SysLog.SetOutput(os.Stdout)
	global.SysLog.SetLevel(logrus.InfoLevel)

	// 测试日志输出
	global.SysLog.Info("日志模块初始化成功")
}
