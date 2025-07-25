package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *ApplicationApi) Role(ctx *gin.Context) {
	DataRole, err := app.Function.Role.Role(ctx)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error while get data"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Role": DataRole,
	})
}
