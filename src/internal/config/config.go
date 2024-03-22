package config

import (
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type config struct {
	DBHost         string `mapstructure:"MYSQL_HOST"`
	DBUserName     string `mapstructure:"MYSQL_USER"`
	DBUserPassword string `mapstructure:"MYSQL_PASSWORD"`
	DBName         string `mapstructure:"MYSQL_DB_NAME"`
	DBPort         string `mapstructure:"MYSQL_PORT"`

	ServerPort string `mapstructure:"SERVER_PORT"`
	ServerHost string `mapstructure:"SERVER_HOST"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`

	AccessTokenExpiresIn  time.Duration
	RefreshTokenExpiresIn time.Duration

	AccessTokenPrivateKey  string `mapstructure:"ACCESS_TOKEN_PRIVATE_KEY"`
	AccessTokenPublicKey   string `mapstructure:"ACCESS_TOKEN_PUBLIC_KEY"`
	RefreshTokenPrivateKey string `mapstructure:"REFRESH_TOKEN_PRIVATE_KEY"`
	RefreshTokenPublicKey  string `mapstructure:"REFRESH_TOKEN_PUBLIC_KEY"`

	AccessTokenMaxAge  int `mapstructure:"ACCESS_TOKEN_MAXAGE"`
	RefreshTokenMaxAge int `mapstructure:"REFRESH_TOKEN_MAXAGE"`
}

func parseConfig() (Config *config, err error) {

	err = godotenv.Load("./config.env")
	if err != nil {
		return nil, err
	}

	var c config

	c.DBHost = os.Getenv("MYSQL_HOST")
	c.DBUserName = os.Getenv("MYSQL_USER")
	c.DBUserPassword = os.Getenv("MYSQL_PASSWORD")
	c.DBName = os.Getenv("MYSQL_DB_NAME")
	c.DBPort = os.Getenv("MYSQL_PORT")

	c.ServerPort = os.Getenv("SERVER_PORT")

	duration, err := time.ParseDuration(os.Getenv("ACCESS_TOKEN_EXPIRED_IN"))
	if err != nil {
		return nil, err
	}

	c.AccessTokenExpiresIn = duration

	duration, err = time.ParseDuration(os.Getenv("REFRESH_TOKEN_EXPIRED_IN"))
	if err != nil {
		return nil, err
	}

	c.RefreshTokenExpiresIn = duration

	c.AccessTokenPrivateKey = os.Getenv("ACCESS_TOKEN_PRIVATE_KEY")
	c.AccessTokenPublicKey = os.Getenv("ACCESS_TOKEN_PUBLIC_KEY")

	c.RefreshTokenPrivateKey = os.Getenv("REFRESH_TOKEN_PRIVATE_KEY")
	c.RefreshTokenPublicKey = os.Getenv("REFRESH_TOKEN_PUBLIC_KEY")

	t, _ := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRED_IN"))
	c.AccessTokenMaxAge = t

	t, _ = strconv.Atoi(os.Getenv("REFRESH_TOKEN_MAXAGE"))
	c.RefreshTokenMaxAge = t

	return &c, err
}

var Config *config

func LoadConfig() *config {
	if Config == nil {
		var err error
		Config, err := parseConfig()
		if err != nil {
			panic(err)
		}
		return Config
	}

	return Config
}
