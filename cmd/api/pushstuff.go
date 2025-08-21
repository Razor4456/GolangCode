package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Razor4456/FoundationBackEnd/internal/store"
	"github.com/gin-gonic/gin"
)

func (app *ApplicationApi) GetDataStuff(ctx *gin.Context) {

	err := app.Function.Stuff.GetDataStuff(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error when get data"})
		return
	}

	// ctx.JSON(http.StatusOK, gin.H{"Stuff": datastuffs})

}

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

	if PayStuff.Namabarang == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Nama barang cannot empty"})
		return
	}

	if PayStuff.Jumlahbarang <= 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Quantity must be greater then 0 "})
		return
	}

	if PayStuff.Harga <= 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Price must be greater then 0"})
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

	if len(PayStuffDelete.Id) == 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Must Input Number"})
		return
	}

	for _, id := range PayStuffDelete.Id {
		if id <= 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{"Message": "Id must be greater than zero "})
			return
		}
	}

	Deleted, err := app.Function.Stuff.DeleteStuff(ctx, PayStuffDelete.Id)

	if err != nil {
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
		"message": "Successfuly deleted item",
		"Data":    Deleted,
	})

}

type StuffEdit struct {
	Id           int64  `json:"id"`
	Namabarang   string `json:"nama_barang"`
	Jumlahbarang int64  `json:"jumlah_barang"`
	Harga        int64  `json:"harga"`
}

func (app *ApplicationApi) EditStuff(ctx *gin.Context) {
	var PayloadEdit StuffEdit

	urlid := ctx.Param("id")

	if urlid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id Cannot Be Null"})
		return
	}

	id, err := strconv.ParseInt(urlid, 10, 64)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Id Format"})
		return
	}

	err = ctx.ShouldBindJSON(&PayloadEdit)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error edit stuff"})
		return
	}

	if id <= 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "id cannot be empty"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if PayloadEdit.Namabarang == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Nama barang cannot be empty"})
		return
	}
	if PayloadEdit.Jumlahbarang <= 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Jumlah barang must be greater then 0"})
		return
	}
	if PayloadEdit.Harga <= 0 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Harga barang must be greater then 0"})
		return
	}

	PayEdit := &store.PostStuff{
		Id:           id,
		Namabarang:   PayloadEdit.Namabarang,
		Jumlahbarang: PayloadEdit.Jumlahbarang,
		Harga:        PayloadEdit.Harga,
	}

	err = app.Function.Stuff.EditStuff(ctx, PayEdit)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Payload Edit Stuff Error"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message":   "Successfuly edit stuff",
		"Data Edit": PayEdit,
	})
}

func (app *ApplicationApi) GetByidStuff(ctx *gin.Context) {
	urlid := ctx.Param("id")

	if urlid == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Id Cannot Be Null"})
		return
	}

	id, err := strconv.ParseInt(urlid, 10, 64)
	if err != nil || id <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid Id Format"})
		return
	}

	// payloadStuffId := StuffGetById{
	// 	Id: StuffbyId.Id,
	// }

	DataStuff, err := app.Function.Stuff.GetByidStuff(ctx, id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error": "Error when sent payload"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{

		"Data": DataStuff,
	})
}
