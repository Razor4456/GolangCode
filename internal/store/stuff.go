package store

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type PostStuff struct {
	Id           int64  `json:"id"`
	Namabarang   string `json:"nama_barang"`
	Jumlahbarang int64  `json:"jumlah_barang"`
	Harga        int64  `json:"harga"`
	CreatedAt    string `json:"created_at"`
}

type StuffApi struct {
	db *sql.DB
}

func (f *StuffApi) CreateStuff(ctx *gin.Context, post *PostStuff) error {

	return nil

}
