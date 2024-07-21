package config

import (
	"os"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

func New[T any]() (T, error) {
	var cfg T

	if envFilePath, ok := os.LookupEnv("ENV_FILE_PATH"); ok {
		if err := godotenv.Load(envFilePath); err != nil {
			return cfg, err
		}
	}

	if err := env.Parse(&cfg); err != nil {
		return cfg, err
	}

	if ok, err := govalidator.ValidateStruct(cfg); !ok {
		return cfg, err
	}

	return cfg, nil
}

type Config struct {
	App        App
	HTTPServer HTTPServer `envPrefix:"HTTP_"`
	Storage    Storage    `envPrefix:"MARIADB_"`
}

type App struct {
	Name string `env:"APP_NAME" envDefault:"app"`
}

type HTTPServer struct {
	Host         string        `env:"SERVER_HOST" envDefault:"localhost"`
	Port         int           `env:"SERVER_PORT" envDefault:"8181"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" envDefault:"15s"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" envDefault:"15s"`
}

type Storage struct {
	DSN     string        `env:"DSN" valid:"required"`
	Timeout time.Duration `env:"QUERY_TIMEOUT" envDefault:"5s"`
}
