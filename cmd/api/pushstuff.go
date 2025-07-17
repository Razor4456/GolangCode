package main

import (
	"errors"
	"fmt"
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

	err := ctx.ShouldBindJSON(&PayStuff)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "There was an error"})
		return
	}

	PayloadsStuff := &store.PostStuff{
		Namabarang:   PayStuff.Namabarang,
		Jumlahbarang: PayStuff.Jumlahbarang,
		Harga:        PayStuff.Harga,
	}

	if err := app.Function.Stuff.CreateStuff(ctx, PayloadsStuff); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "Cannot post data stuff"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Message": "Successfuly create stuff",
		"data":    PayloadsStuff})

}

type PayloadStuffDelete struct {
	Id []int64 `json:"id"`
}

func (app *ApplicationApi) DeleteStuff(ctx *gin.Context) {
	var PayStuffDelete PayloadStuffDelete

	err := ctx.ShouldBindJSON(&PayStuffDelete)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "There was an error"})
		return
	}

	if err := app.Function.Stuff.DeleteStuff(ctx, PayStuffDelete.Id); err != nil {
		if errors.Is(err, store.ErrNotFound) {
			ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("stuff with id %d not found", PayStuffDelete.Id)})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "failed to delete stuff",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("stuff with id: %d deleted successfuly", PayStuffDelete.Id),
	})

}
