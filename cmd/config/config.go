package config

import (
	"reflect"

	"github.com/balabanovds/system-monitor/internal/models"
	"github.com/caarlos0/env/v6"
)

type Config struct {
	App    AppConfig
	Server ServerConfig
	Log    LogConfig
}

type AppConfig struct {
	IntervalSeconds     int                 `env:"TICK_INTERVAL" envDefault:"1"`
	MaxMeasurementHours int                 `env:"MAX_MEASURE_INTERVAL" envDefault:"1"`
	Parsers             []models.ParserType `env:"PARSERS" envSeparator:":" envDefault:"load_avg:cpu:io"`
}

type ServerConfig struct {
	Host     string `env:"HOST" envDefault:"0.0.0.0"`
	GRPCPort int    `env:"GRPC_PORT" envDefault:"9000"`
	HTTPPort int    `env:"HTTP_PORT" envDefault:"9001"`
}

type LogConfig struct {
	Level      string `env:"LOG_LEVEL" envDefault:"info"`
	Production bool   `env:"PRODUCTION" envDefault:"false"`
}

func Parse() (*Config, error) {
	cfg := new(Config)

	return cfg, env.ParseWithFuncs(cfg, map[reflect.Type]env.ParserFunc{
		reflect.TypeOf(models.Undef): typeParser,
	})
}

func typeParser(str string) (interface{}, error) {
	var p models.ParserType

	return p.Value(str), nil
}
