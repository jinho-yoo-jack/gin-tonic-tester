package config

import (
	"fmt"
	"ginTonicProject/utils"
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

type Config struct {
	Environment   string `mapstructure:"ENVIRONMENT"`
	MysqlHost     string `mapstructure:"MYSQL_HOST"`
	MysqlUserName string `mapstructure:"MYSQL_USERNAME"`
	MysqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	MysqlDbname   string `mapstructure:"MYSQL_DBNAME"`
}

func MustLoadConfig() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg
}

func LoadConfig() (*Config, error) {
	projectRoot, err := utils.FindProjectRootFromGoMod()
	if err != nil {
		return nil, err
	}

	viper.SetConfigFile(filepath.Join(projectRoot, "app.env"))
	viper.SetConfigType("env")
	viper.AutomaticEnv() // 환경변수 우선

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("config load error: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("config unmarshal error: %w", err)
	}
	return &cfg, nil
}
