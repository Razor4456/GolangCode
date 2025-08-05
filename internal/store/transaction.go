package store

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transaction struct {
	Id     int64        `json:"id"`
	UserId int64        `json:"user_id"`
	Barang []TransStuff `json:"stuff"`
}

type TransStuff struct {
	IdBarang     int64  `json:"id_barang"`
	Namabarang   string `json:"nama_barang"`
	Jumlahbarang int64  `json:"jumlahbarang"`
	Harga        int64  `json:"harga"`
}

type TransactionAPI struct {
	db *sql.DB
}

func (f *TransactionAPI) Cart(ctx *gin.Context, trx *Transaction) error {
	tx, err := f.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	var TransactionDatas []TransStuff

	for _, item := range trx.Barang {
		var BarangId int64
		var NamaBarang string
		var StockBarang int64
		var HargaBarang int64

		query := `SELECT id, nama_barang, jumlah_barang, harga FROM stuff WHERE id = $1`

		err := tx.QueryRowContext(
			ctx,
			query,
			item.IdBarang,
		).Scan(
			&BarangId,
			&NamaBarang,
			&StockBarang,
			&HargaBarang,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No Data Found"})
				return err
			}
			return fmt.Errorf("Eidt:%s", err)
		}

		if item.Jumlahbarang > StockBarang {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Stock tidak mencukupi / Stock Habis"})
			tx.Rollback()
			return nil
		}

		queryupdate := `UPDATE stuff SET jumlah_barang = jumlah_barang - $1 WHERE id = $2`
		_, err = tx.ExecContext(
			ctx,
			queryupdate,
			item.Jumlahbarang,
			item.IdBarang,
		)

		if err != nil {
			tx.Rollback()
			return err
		}

		var DataTransaction TransStuff
		insert := `INSERT INTO transactions (userid, idbarang, nama_barang, jumlah_barang, harga) VALUES ($1, $2, $3, $4, $5) RETURNING (userid, idbarang, nama_barang, jumlah_barang, harga)`
		err = tx.QueryRowContext(
			ctx,
			insert,
			trx.UserId,
			item.IdBarang,
			NamaBarang,
			item.Jumlahbarang,
			HargaBarang,
		).Scan(
			&DataTransaction.IdBarang,
			&DataTransaction.Namabarang,
			&DataTransaction.Jumlahbarang,
			&DataTransaction.Harga,
		)

		TransactionDatas = append(TransactionDatas, DataTransaction)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()

	if err != nil {
		return err
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Transaksi Berhasil",
		"data": TransactionDatas,
	})

	return nil
}
