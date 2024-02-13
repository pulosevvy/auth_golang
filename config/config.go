package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

const (
	configPath = "config/local.yaml"
)

type (
	Config struct {
		Env        string `yaml:"env" env-default:"local" env-required:"true"`
		HTTPServer `yaml:"http_server"`
		PG         `yaml:"pg"`
		Mongo      `yaml:"mongo"`
		Token      `yaml:"token"`
	}

	HTTPServer struct {
		Address     string        `yaml:"address" env-default:"localhost:8080"`
		Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
		IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
	}

	PG struct {
		Host     string `yaml:"host" env-required:"true"`
		Port     string `yaml:"port" env-required:"true"`
		Username string `yaml:"username" env-required:"true"`
		Password string `yaml:"password" env-required:"true"`
		DBName   string `yaml:"db_name" env-required:"true"`
		SslMode  string `yaml:"ssl_mode" env-required:"true"`
	}

	Mongo struct {
		Host     string `yaml:"host" env-required:"true"`
		Port     string `yaml:"port" env-required:"true"`
		Username string `yaml:"username" env-required:"true"`
		Password string `yaml:"password" env-required:"true"`
		DBName   string `yaml:"db_name" env-required:"true"`
	}

	Token struct {
		AccessTokenTTL  time.Duration `yaml:"accessTokenTTL" env-required:"true"`
		RefreshTokenTTL string        `yaml:"refreshTokenTTL" env-required:"true"`
		PrivateKey      string        `yaml:"privateKey" env-required:"true"`
	}
)

func MustLoad() (*Config, error) {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exists: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg, nil
}
