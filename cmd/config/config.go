package config

import (
	"encoding/json"
	"os"

	"github.com/balabanovds/smonitor/internal/api"
	"github.com/balabanovds/smonitor/internal/app"
)

type Config struct {
	App  app.Config `json:"app"`
	Grpc api.Config `json:"grpc"`
}

func Unmarshal(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	var c *Config

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
