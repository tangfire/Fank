package configs

import (
	"github.com/spf13/viper"
	"log"
)

// DatabaseConfig 存储数据库相关配置
type DatabaseConfig struct {
	DBName     string `mapstructure:"DB_NAME"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PSW"`
}

type Config struct {
	DBConfig DatabaseConfig `mapstructure:"database"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("./configs/config.yml")
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
