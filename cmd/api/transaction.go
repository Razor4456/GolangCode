package main

import (
	"net/http"

	"github.com/Razor4456/FoundationBackEnd/internal/store"
	"github.com/gin-gonic/gin"
)

type TransactionPayload struct {
	UserId int64           `json:"user_id"`
	Barang []BarangPayload `json:"stuff"`
}

type BarangPayload struct {
	IdBarang     int64 `json:"id_barang"`
	Jumlahbarang int64 `json:"jumlah_barang"`
}

func (app *ApplicationApi) Cart(ctx *gin.Context) {
	var PayTransaction TransactionPayload
	err := ctx.ShouldBindJSON(&PayTransaction)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "There was an error input payload"})
		return
	}

	PayloadTransaction := &store.Transaction{
		UserId: PayTransaction.UserId,
		Barang: make([]store.TransStuff, 0),
	}

	for _, ProsesTrx := range PayTransaction.Barang {

		item := store.TransStuff{
			IdBarang:     ProsesTrx.IdBarang,
			Jumlahbarang: ProsesTrx.Jumlahbarang,
		}

		PayloadTransaction.Barang = append(PayloadTransaction.Barang, item)
	}

	err = app.Function.Transaction.Cart(ctx, PayloadTransaction)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "There was an error when doing transaction"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message":          "Transaction successfuly",
		// "Transaction Data": PayTransaction,
	})
}
