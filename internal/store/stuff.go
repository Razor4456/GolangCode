package store

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type PostStuff struct {
	Id           int64  `json:"id"`
	Namabarang   string `json:"nama_barang"`
	Jumlahbarang int64  `json:"jumlah_barang"`
	Harga        int64  `json:"harga"`
	CreatedAt    string `json:"created_at"`
}

type DeletedStuff struct {
	Namabarang string `json:"nama_barang"`
}

type StuffApi struct {
	db *sql.DB
}

func (f *StuffApi) GetDataStuff(ctx *gin.Context) ([]PostStuff, error) {
	query := `SELECT * FROM stuff`

	stuffrows, err := f.db.Query(query)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "There was an error on GetDataStuff"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return nil, nil
	}

	defer stuffrows.Close()

	var DataStuff []PostStuff

	for stuffrows.Next() {
		var datastuffrows PostStuff
		err := stuffrows.Scan(
			&datastuffrows.Id,
			&datastuffrows.Namabarang,
			&datastuffrows.Jumlahbarang,
			&datastuffrows.Harga,
			&datastuffrows.CreatedAt,
		)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Ther was an error when get datastuff"})
			ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
			return nil, nil

		}
		DataStuff = append(DataStuff, datastuffrows)
	}

	if err = stuffrows.Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "There was an error on GetDataStuff"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		return nil, nil
	}

	return DataStuff, nil

}

func (f *StuffApi) CreateStuff(ctx *gin.Context, poststuff *PostStuff) error {
	query := `INSERT INTO stuff(nama_barang, jumlah_barang, harga)
	VALUES($1,$2,$3) RETURNING id, created_at`

	err := f.db.QueryRowContext(
		ctx,
		query,
		poststuff.Namabarang,
		poststuff.Jumlahbarang,
		poststuff.Harga,
	).Scan(
		&poststuff.Id,
		&poststuff.CreatedAt,
	)

	if err != nil {
		log.Printf("Failed To Insert Data Error: %s, data: %v", err, poststuff)
		return fmt.Errorf("failed to make stuff error:%s", err)
	}

	return nil
}

func (f *StuffApi) DeleteStuff(ctx *gin.Context, StuffId []int64) ([]DeletedStuff, error) {
	query := `DELETE FROM stuff WHERE id = any($1)
	RETURNING nama_barang`

	result, err := f.db.QueryContext(ctx, query, pq.Array(StuffId))

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error:": "There was an error on DeleteStuff"})
		return nil, fmt.Errorf("failed to delete: %w", err)
	}

	defer result.Close()

	var DeletedBarang []DeletedStuff

	for result.Next() {
		var PosStuff DeletedStuff
		if err := result.Scan(&PosStuff.Namabarang); err != nil {
			return nil, fmt.Errorf("failed to return name: %w", err)
		}
		DeletedBarang = append(DeletedBarang, PosStuff)
	}

	if err != nil {
		return nil, fmt.Errorf("rows error :%w", err)
	}

	return DeletedBarang, nil
}
