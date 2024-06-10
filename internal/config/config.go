package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string     `yaml:"env" env-default:"local"` // Прочитаь про struct tag
	StoragePath string     `yaml:"storage_path" rnv-required:"true"`
	HTTPServer  HTTPServer `yaml:"http_server"`
}

type HTTPServer struct {
	Host        string        `yaml:"host" env-default:"localhost"`
	Port        string        `yaml:"port" env-default:"8080"`
	Timeout     time.Duration `yaml:"timout" env-default:"5s"`
	TdleTimeout time.Duration `yaml:"tdle_timeout" env-default:"60s"`
}

func MustLoad() *Config {

	// Провекрка существования .env
	if err := godotenv.Load(); err != nil {
		log.Print(".env file found")
	}

	configPath, exists := os.LookupEnv("CONFIG_PATH")
	if !exists {
		log.Fatal("CONFIG_PATH is not found")
	}

	// Провека на существования файла
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
