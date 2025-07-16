package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *ApplicationApi) Servererror(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Status internal server error"})
}
