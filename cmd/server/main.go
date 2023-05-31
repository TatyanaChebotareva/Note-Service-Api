package main

import (
	"context"
	"flag"
	"log"

	app "github.com/TatyanaChebotareva/Note-Service-Api/internal/app/api"
	_ "github.com/jackc/pgx/stdlib"
)

var pathConfig string

func init() {
	flag.StringVar(&pathConfig, "config", "./config/config.json", "Path to configuration file")
}

func main() {
	flag.Parse()

	ctx := context.Background()
	a, err := app.NewApp(ctx, pathConfig)
	if err != nil {
		log.Fatalf("failed to create app: %s", err.Error())
	}

	err = a.Run()
	if err != nil {
		log.Fatalf("failed to create app: %s", err.Error())
	}
}
