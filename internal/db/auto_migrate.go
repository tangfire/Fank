package db

import (
	"Fank/internal/global"
	"Fank/internal/model"
	"log"
)

func autoMigrate() {
	if global.DB == nil {
		log.Fatal("数据库初始化失败，无法执行自动迁移...")
	}

	err := global.DB.AutoMigrate(model.GetAllModels()...)

	if err != nil {
		log.Fatalf("autoMigrate err: %v", err)
	}

	log.Println("autoMigrate success")
	global.SysLog.Infof("autoMigrate success")
}
