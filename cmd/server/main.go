package main

import (
	"context"
	"flag"
	"os"

	"github.com/joho/godotenv"

	"github.com/huey-emma/cms/db"
	"github.com/huey-emma/cms/internal/router"
	"github.com/huey-emma/cms/internal/utils/jsonlog"
)

type config struct {
	port string
	db   struct {
		uri string
		maxopen,
		maxidle,
		idletime int
	}
}

func main() {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)
	godotenv.Load()

	conf := config{}

	flag.StringVar(&conf.port, "port", os.Getenv("PORT"), "app port")
	flag.StringVar(&conf.db.uri, "postgres_uri", os.Getenv("POSTGRES_URI"), "postgres uri")
	flag.IntVar(&conf.db.maxidle, "max_idle", 25, "max idle connections")
	flag.IntVar(&conf.db.maxopen, "max_open", 25, "max open connections")
	flag.IntVar(&conf.db.idletime, "idle_time", 15, "idle connections max timeout")

	flag.Parse()

	database, err := db.New(conf.db.uri, conf.db.maxopen, conf.db.maxidle, conf.db.idletime)
	if err != nil {
		logger.PrintFatal(err, nil)
	}

	logger.PrintInfo("database connected!", nil)

	if err := database.PingContext(context.Background()); err != nil {
		logger.PrintFatal(err, nil)
	}

	logger.PrintInfo("successfully pinged the database", nil)

	if err := database.MigrateDb(); err != nil {
		logger.PrintFatal(err, nil)
	}

	logger.PrintInfo("database migrations done!", nil)

	r := router.New(conf.port, logger, database)

	if err := r.Use().Run(); err != nil {
		logger.PrintFatal(err, nil)
	}
}
