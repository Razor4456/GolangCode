package store

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type FunctionStore struct {
	Stuff interface {
		CreateStuff(*gin.Context, *PostStuff) error
	}
}

func FunctionStorage(db *sql.DB) FunctionStore {
	return FunctionStore{
		Stuff: &StuffApi{db},
	}
}
