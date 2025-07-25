package main

import (
	"net/http"

	"github.com/Razor4456/FoundationBackEnd/internal/store"
	"github.com/Razor4456/FoundationBackEnd/utils"
	"github.com/gin-gonic/gin"
)

type LoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (app *ApplicationApi) Login(ctx *gin.Context) {
	var LoginPayload LoginPayload

	err := ctx.ShouldBindJSON(&LoginPayload)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error with payload Login, LoginPayload"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	payloadlogin := &store.UsersLogin{
		Username: LoginPayload.Username,
		Password: LoginPayload.Password,
	}

	err = app.Function.Users.Login(ctx, payloadlogin)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error when login"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Login Success"})

}

type PayloadUsers struct {
	Email      string `json:"email"`
	Username   string `json:"username"`
	Name       string `json:"name"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	VerifLogin string `json:"veriflogin"`
}

func (app *ApplicationApi) CreateUsers(ctx *gin.Context) {
	var Paypostusers PayloadUsers

	err := ctx.ShouldBindJSON(&Paypostusers)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "There was an error with payload"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	passhash, err := utils.HashPassword(Paypostusers.Password)

	if err != nil {
		ctx.JSON(http.StatusPreconditionFailed, gin.H{"Message": "There was an error when parshing hash password"})
		return
	}

	InputUsers := &store.PostUsers{
		Email:      Paypostusers.Email,
		Username:   Paypostusers.Username,
		Name:       Paypostusers.Name,
		Password:   passhash,
		Role:       Paypostusers.Role,
		VerifLogin: "False",
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
