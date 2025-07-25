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

	tokenz, err := app.Function.Users.Login(ctx, payloadlogin)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "There was an error when login"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"Message":  "Login Success",
		"User_id":  tokenz.Id,
		"Username": tokenz.Name,
		"Token":    tokenz.Token,
	})

}

type PayloadLogout struct {
	Username string `json:"username"`
}

func (app *ApplicationApi) LogOut(ctx *gin.Context) {
	var PayloadLogout PayloadLogout

	err := ctx.ShouldBindJSON(&PayloadLogout)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": "There was an error in Log out Function"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}

	Paylogout := &store.StoreLogout{
		Username: PayloadLogout.Username,
	}

	err = app.Function.Users.Logout(ctx, Paylogout)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Logout function error"})
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Log Out successfuly"})

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
