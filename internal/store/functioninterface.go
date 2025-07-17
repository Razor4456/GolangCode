package store

import (
	"database/sql"
	"errors"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	ErrNotFound = errors.New("record not found")
)

type FunctionStore struct {
	Stuff interface {
		CreateStuff(*gin.Context, *PostStuff) error
		DeleteStuff(*gin.Context, []int64) ([]DeletedStuff, error)
	}
}

func FunctionStorage(db *sql.DB) FunctionStore {
	return FunctionStore{
		Stuff: &StuffApi{db},
	}
}
