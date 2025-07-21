package main

import (
	"net/http"

	"github.com/Razor4456/FoundationBackEnd/internal/store"
	"github.com/Razor4456/FoundationBackEnd/utils"
	"github.com/gin-gonic/gin"
)

type PayloadUsers struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

func (app *ApplicationApi) CreateUsers(ctx *gin.Context) {
	var Paypostusers PayloadUsers

	err := ctx.ShouldBindJSON(&Paypostusers)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "There was an error with payload"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	Passhash, err := utils.HashPassword(Paypostusers.Password)

	InputUsers := &store.PostUsers{
		Email:    Paypostusers.Email,
		Username: Paypostusers.Username,
		Name:     Paypostusers.Name,
		Password: Passhash,
		Role:     Paypostusers.Role,
	}

	err = app.Function.Users.CreateUsers(ctx, InputUsers)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": "There was an error when create users"})
		ctx.JSON(http.StatusBadRequest, gin.H{"Error:": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Message": "user successfuly created",
		"data":    InputUsers,
	})

}
