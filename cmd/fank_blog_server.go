package cmd

import (
	"Fank/configs"
	"Fank/internal/db"
	"github.com/labstack/echo/v4"
	"log"
)

func Start() {
	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatalf("程序启动时加载配置失败: %v", err)
		return
	}

	// 初始化 echo 实例
	app := echo.New()
	//app.HideBanner = true

	db.New(config)

	//app.Logger.Fatal(app.Start(fmt.Sprintf()))

}
