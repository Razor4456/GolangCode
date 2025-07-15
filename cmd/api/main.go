package main

import "github.com/Razor4456/FoundationBackEnd/internal/env"

func main() {
	cfg := Config{
		Addr: env.GetString("ADDR", ":8080"),
		Db: Dbconfig{
			Addr:         env.GetString("DB_ADDR", "postgres://postgres:raxon789@localhost/?sslmode=disable"),
			MaxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			MaxIdleConss: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			MaxIdleTime:  env.GetString("DB_MAX_IDLE_CONNS", "15m"),
		},
	}
}
