package middlewares

import (
	"net/http"
	"strings"

	"github.com/Razor4456/FoundationBackEnd/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")

	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	tokenstring := strings.TrimPrefix(token, "Bearer")

	if tokenstring == token {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token Format. Expected 'Bearer <token>'"})
		return
	}

	claims, err := utils.VerifToken(tokenstring)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error":   "invalid Token Details",
			"details": err.Error()})
		return
	}

	ctx.Set("User_id", claims.UserId)
	ctx.Set("username", claims.Username)

	ctx.Next()
}
