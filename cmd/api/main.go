package main

import (
	"log"

	"github.com/Razor4456/FoundationBackEnd/internal/db"
	"github.com/Razor4456/FoundationBackEnd/internal/env"
	"github.com/Razor4456/FoundationBackEnd/internal/store"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := Config{
		Addr: env.GetString("ADDR", ":8080"),
		Db: Dbconfig{
			Addr:         env.GetString("DB_ADDR", "postgres://postgres:raxon789@localhost/newdatabase?sslmode=disable"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConss: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetString("DB_MAX_IDLE_CONNS", "15m"),
		},
		Env: env.GetString("ENV", "Development"),
	}

	db, err := db.Database(cfg.Db.Addr,
		cfg.Db.MaxOpenConns,
		cfg.Db.MaxIdleConss,
		cfg.Db.MaxIdleTime)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	log.Printf("Database Successfully Connect")

	store := store.FunctionStorage(db)

	app := &ApplicationApi{
		Config:   cfg,
		Function: store,
	}

	ginserver := gin.Default()

	app.ServerRoute(ginserver)

	log.Printf("server Has Starting on, %s", cfg.Addr)

	if err := ginserver.Run(cfg.Addr); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
