package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HTTPServer HTTPServer
}

type HTTPServer struct {
	Addr string `envconfig:"HTTP_ADDR" default:"0.0.0.0:8080"`
}

func Load() (*Config, error) {
	c := Config{}
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
