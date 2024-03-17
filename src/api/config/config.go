package config

import (
	"crypto/rand"
	"crypto/rsa"
)

var config *configuration

func initConfig() *configuration {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	return &configuration{
		privateKey: key,
	}
}

func GetConfig() *configuration {
	if config == nil {
		config = initConfig()
		return config
	}
	return config
}

type configuration struct {
	privateKey *rsa.PrivateKey
}

func (c *configuration) GetPrivateKey() rsa.PrivateKey {
	return *c.privateKey
}
