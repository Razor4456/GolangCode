package utils

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

// var secretKey = "thisissecretkey"

func GenerateToken(userId int64, Username string) (string, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		return "", errors.New("JWT SECRET not found")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":   userId,
		"username": Username,
		"exp":      time.Now().Add(20 * time.Minute).Unix(),
	})

	return token.SignedString([]byte(secretKey))
}
