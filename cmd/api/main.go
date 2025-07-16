package main

import "github.com/gin-gonic/gin"

func (app *ApplicationApi) main() {
	DbConnection()

	ginserver := gin.Default()

	app.ServerRoute(ginserver)

}
