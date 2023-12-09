package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
	"time"
)

type (
	Config struct {
		AppCfg    `yaml:"app"`
		ServerCfg `yaml:"server"`
		PG        `yaml:"pg"`
	}

	AppCfg struct {
		Name string `yaml:"name"`
	}

	ServerCfg struct {
		Port            string        `yaml:"port"`
		ReadTimeout     time.Duration `yaml:"readTimeout"`
		WriteTimeout    time.Duration `yaml:"writeTimeout"`
		ShutdownTimeout time.Duration `yaml:"shutdownTimeout"`
	}

	PG struct {
		ConnURI string `yaml:"connURI"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, errors.Wrap(err, "NewConfig: the config files cannot be read, an error has returned")
	}

	return cfg, nil
}
