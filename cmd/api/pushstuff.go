package main

import (
	"net/http"

	"github.com/Razor4456/FoundationBackEnd/internal/store"
	"github.com/gin-gonic/gin"
)

type PayloadStuff struct {
	Namabarang   string `json:"nama_barang"`
	Jumlahbarang int64  `json:"jumlah_barang"`
	Harga        int64  `json:"harga"`
}

func (app *ApplicationApi) CreateStuff(ctx *gin.Context) {
	var PayStuff PayloadStuff

	err := ctx.ShouldBind(&PayStuff)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "There was an error"})
		return
	}

	PayloadsStuff := &store.PostStuff{
		Namabarang:   PayStuff.Namabarang,
		Jumlahbarang: PayStuff.Jumlahbarang,
		Harga:        PayStuff.Harga,
	}

}
