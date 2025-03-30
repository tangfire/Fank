package configs

import (
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	AppName string `mapstructure:"APP_NAME"`
	AppHost string `mapstructure:"APP_HOST"`
	AppPort string `mapstructure:"APP_PORT"`
}

// DatabaseConfig 存储数据库相关配置
type DatabaseConfig struct {
	DBName     string `mapstructure:"DB_NAME"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PSW"`
}

type LogConfig struct {
	LogFilePath     string `mapstructureL:"LOG_FILE_PATH"`
	LogFileName     string `mapstructureL:"LOG_FILE_NAME"`
	LogTimestampFmt string `mapstructureL:"LOG_TIMESTAMP_FMT"`
	LogMaxAge       int    `mapstructureL:"LOG_MAX_AGE"`
	LogRotationTime int    `mapstructureL:"LOG_ROTATION_TIME"`
	LogLevel        string `mapstructureL:"LOG_LEVEL"`
}

type Config struct {
	AppConfig AppConfig      `mapstructure:"app"`
	DBConfig  DatabaseConfig `mapstructure:"database"`
	LogConfig LogConfig      `mapstructure:"log"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("./configs/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %s", err)
		return nil, err
	}

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
		return nil, err
	}

	return &config, nil

}
