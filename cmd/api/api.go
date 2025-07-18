package main

import (
	"github.com/Razor4456/FoundationBackEnd/internal/store"
	"github.com/gin-gonic/gin"
)

type ApplicationApi struct {
	Config   Config
	Function store.FunctionStore
}

type Config struct {
	Addr string
	Db   Dbconfig
	Env  string
}

type Dbconfig struct {
	Addr         string
	MaxOpenConns int
	MaxIdleConss int
	MaxIdleTime  string
}

func (app *ApplicationApi) ServerRoute(route *gin.Engine) {

	Found := route.Group("/FoundationV1")
	{
		Found.GET("/GetStuff", app.GetDataStuff)
		Found.POST("/AddStuff", app.CreateStuff)
		Found.DELETE("/DeleteStuff", app.DeleteStuff)
	}

}
