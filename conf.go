package main

import (
	"github.com/envconfig"
)

const (
	AppName = "redis_test"
)

//Конфигурация Redis
type Config struct {
	DBType    string `envconfig:"db_type" default:"redis"`
	RedisHost string `envconfig:"redis_host" default:"localhost:6379"`
	RedisPass string `envconfig:"redis_pass" default:""` // default to no password
	RedisDB   int64  `envconfig:"redis_db" default:"0"`  // default to the redis default DB
}

// Инициализирует и возвращает конфигурацию или ошибку, если произошла
func GetConfig() (*Config, error) {
	conf := new(Config)
	if err := envconfig.Process(AppName, conf); err != nil {
		return nil, err
	}
	return conf, nil
}
