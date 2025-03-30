package db

import (
	"Fank/internal/global"
	"log"
)

func autoMigrate() {
	if global.DB == nil {
		log.Fatal("数据库初始化失败，无法执行自动迁移...")
	}

	global.DB.AutoMigrate()
}
