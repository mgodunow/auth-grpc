package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string        `yaml:"env" env:"ENV"`
	StoragePath string        `yaml:"storage_path" env:"STORAGE_PATH"`
	TokenTTL    time.Duration `yaml:"token_ttl" env:"TOKEN_TTL"`
	GRPC        GRPCConfig
}

type GRPCConfig struct {
	Port    int           `yaml:"port" env:"PORT"`
	Timeout time.Duration `yaml:"timeout" env:"TIMEOUT"`
}

func MustLoad() *Config {
	path := fetchConfigPath()
	if path == "" {
		panic("empty path to config file: ")
	}
	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file does not exist: " + path)
	}
	var config Config
	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("failed to read config: " + err.Error())
	}
	var grpc GRPCConfig
	if err := cleanenv.ReadConfig(path, &grpc); err != nil {
		panic("failed to read GRPC config: " + err.Error())
	}
	config.GRPC = grpc
	return &config
}

func fetchConfigPath() string {
	result := ""

	flag.StringVar(&result, "config", "", "path to config file")
	flag.Parse()
	if result == "" {
		result = os.Getenv("CONFIG_PATH")
	}
	return result
}
