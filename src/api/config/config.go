package config

import (
	"time"

	"github.com/spf13/viper"
)

type config struct {
	DBHost         string `mapstructure:"MYSQL_HOST"`
	DBUserName     string `mapstructure:"MYSQL_USER"`
	DBUserPassword string `mapstructure:"MYSQL_PASSWORD"`
	DBName         string `mapstructure:"MYSQL_DB"`
	DBPort         string `mapstructure:"MYSQL_PORT"`
	ServerPort     string `mapstructure:"PORT"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
	RedisUri     string `mapstructure:"REDIS_URL"`

	AccessTokenPrivateKey  string        `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string        `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string        `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string        `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`
	AccessTokenExpiresIn   time.Duration `mapstructure:"ACCESS_TOKEN_EXPIRED_IN"`
	RefreshTokenExpiresIn  time.Duration `mapstructure:"REFRESH_TOKEN_EXPIRED_IN"`
	AccessTokenMaxAge      int           `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge     int           `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

func loadConfig(path string) (Config *config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigType("env")
	viper.SetConfigName("config")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var c config
	err = viper.Unmarshal(&c)
	return &c, err
}

var Config *config

func GetConfig() config {
	if Config != nil {
		Config, err := loadConfig("/")
		if err != nil {
			panic(err)
		}
		return *Config
	}

	return *Config
}
