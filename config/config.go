package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Environment   string `mapstructure:"ENVIRONMENT"`
	MysqlHost     string `mapstructure:"MYSQL_HOST"`
	MysqlUserName string `mapstructure:"MYSQL_USERNAME"`
	MysqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	MysqlDbname   string `mapstructure:"MYSQL_DBNAME"`
}

func LoadConfig(path string) (config Config, error error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
