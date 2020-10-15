package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/balabanovds/smonitor/internal/api"

	"github.com/balabanovds/smonitor/internal/app"
	"github.com/balabanovds/smonitor/internal/metrics/inmem"

	"github.com/balabanovds/smonitor/cmd/config"
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

	cfg, err := config.Unmarshal(configFile)
	if err != nil {
		log.Fatalln(err)
	}

	srv := api.NewServer(cfg.Grpc)

	ctx, cnc := context.WithCancel(context.Background())
	defer cnc()

	a := app.New(cfg.App, inmem.New())

	go func() {
		log.Fatalln(srv.Serve(*a))
	}()
	<-a.Run(ctx)
}
