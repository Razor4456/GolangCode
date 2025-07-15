package main

import "github.com/Razor4456/FoundationBackEnd/internal/store"

type ApplicationApi struct {
	Config   Config
	Function store.FunctionStore
}

type Config struct {
	Addr string
	Db   Dbconfig
	Env  string
}

type Dbconfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConss int
	MaxIdleTime  string
}
