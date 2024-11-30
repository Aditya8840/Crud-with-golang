package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)


type HttpServer struct {
	Address string `yaml:"address" env-required:"true"`
}

type Config struct {
	Env         string `yaml:"env" env-required:"true"`
	StoragePath string   `yaml:"storage_path" env-required:"true"`
	HttpServer `yaml:"http_server" env-required:"true"`
}

func MustLoad() *Config {
	var configPath string
	
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("config-path", "", "path to config file")
		flag.Parse()
		configPath = *flags

		if configPath == "" {
			log.Fatal("config path is not set")
		}
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("failed to read config: %s", err)
	}

	return &cfg
}
