package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App      AppConfig      `json:"app"`
	DB       DBConfig       `json:"db"`
	Redis    RedisConfig    `json:"redis"`
	RabbitMq RabbitMqConfig `json:"rabbitmq"`
}

type DBConfig struct {
	Dialect      string `json:"dialect"`
	Datasource   string `json:"datasource"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_open_conns"`
}

type AppConfig struct {
	Addr string `json:"addr"`
}

type RedisConfig struct {
	Address string `json:"address"`
}

type RabbitMqConfig struct {
	Url string `json:"url"`
}

var (
	App      AppConfig
	DB       DBConfig
	Redis    RedisConfig
	RabbitMq RabbitMqConfig
)

func SetupConfig() (config *Config) {
	config = &Config{}

	viper.AddConfigPath("config/")
	viper.SetConfigName("config_files")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	_ = viper.Unmarshal(config)

	App = config.App
	DB = config.DB
	Redis = config.Redis
	RabbitMq = config.RabbitMq
	return
}
