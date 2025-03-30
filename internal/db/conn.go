package db

import (
	"Fank/configs"
	"Fank/internal/global"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// New 初始化数据库连接
func New(config *configs.Config) {
	var err error
	global.DB, err = connectToDB(config, config.DBConfig.DBName)
	if err != nil {
		log.Fatalf("connect to DB failed: %v", err)
		return
	}
	log.Printf("connect to [%s] DB success", config.DBConfig.DBName)
	global.SysLog.Infof("connect to DB [%s] success", config.DBConfig.DBName)

	autoMigrate()

}

// connectToDB 数据库连接
func connectToDB(config *configs.Config, dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.DBConfig.DBUser, config.DBConfig.DBPassword, config.DBConfig.DBHost, config.DBConfig.DBPort, dbName)

	mysqlDialector := mysql.Open(dsn)
	return gorm.Open(mysqlDialector, &gorm.Config{})
}
