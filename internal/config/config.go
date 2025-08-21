package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"local"`
	Storage    `yaml:"storage"`
	HttpServer `yaml:"http_server"`
}

type Storage struct {
	User     string `yaml:"user" env-required:"true"`
	Password string `yaml:"pass" env-required:"true"`
	Address  string `yaml:"address" env-default:"0.0.0.0:5438"`
	NameDB   string `yaml:"name_db" env-default:"spot_db"`
}

type HttpServer struct {
	Address     string        `yaml:"address" env-default:"localhost:3030"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// check if file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
