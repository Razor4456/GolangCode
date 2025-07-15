package store

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type PostStuff struct {
}

type StuffApi struct {
	db *sql.DB
}

func (f *StuffApi) CreateStuff(ctx *gin.Context, post *PostStuff) error {

	return nil

}
