package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/balabanovds/system-monitor/cmd/config"
	"github.com/balabanovds/system-monitor/cmd/logger"
	"github.com/balabanovds/system-monitor/internal/api"
	"github.com/balabanovds/system-monitor/internal/app"
	"github.com/balabanovds/system-monitor/internal/storage/inmem"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "json", "./config/config.json", "JSON config file")
	flag.Parse()
}

func main() {
	if configFile == "" {
		flag.Usage()
		os.Exit(1)
	}

	cfg, err := config.Parse()
	handleErr(err)

	zapLogger, err := logger.New(cfg.Log.Level, cfg.Log.Production)
	handleErr(err)

	srv := api.NewServer(cfg.Server, zapLogger)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a := app.New(cfg.App, inmem.New(zapLogger), zapLogger)

	go func() {
		log.Fatalln(srv.Serve(*a))
	}()
	<-a.Run(ctx)
}

func handleErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
