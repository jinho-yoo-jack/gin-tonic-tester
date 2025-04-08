package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
)

type Config struct {
	Environment     string `mapstructure:"ENVIRONMENT"`
	MysqlHost       string `mapstructure:"MYSQL_HOST"`
	MysqlUserName   string `mapstructure:"MYSQL_USERNAME"`
	MysqlPassword   string `mapstructure:"MYSQL_PASSWORD"`
	MysqlDbname     string `mapstructure:"MYSQL_DBNAME"`
	SecretKey       string `mapstructure:"SECRET_KEY"`
	AccessTokenExp  string `mapstructure:"ACCESS_TOKEN_EXP"`
	RefreshTokenExp string `mapstructure:"REFRESH_TOKEN_EXP"`
}

func MustLoadConfig() *Config {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	return cfg
}

func LoadConfig() (*Config, error) {
	projectRoot, err := findProjectRootFromGoMod()
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

func findProjectRootFromGoMod() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found")
		}
		dir = parent
	}
}
