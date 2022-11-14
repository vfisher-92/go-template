package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

const EnvPrefix = "SERVICE_"

type Config struct {
	ServerConfig

	DB    DatabaseConfig `env:"DB"      json:"db"`
	Redis RedisConfig    `env:"REDIS"   json:"redis"`

	DevMode bool `env:"DEV_MODE"                 json:"dev_mode"`
}

var config *Config

func InitConfig() {
	config = &Config{}

	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatal(err)
	}

	opts := env.Options{
		Prefix: EnvPrefix,
	}

	if err := env.Parse(config, opts); err != nil {
		log.Fatal(err)
	}
}

func GetInstance() *Config {
	return config
}
