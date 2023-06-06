package utils

import (
	"errors"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte("Secret-Token-Tecnica")

// GenerateToken generates a new JWT token
func GenerateToken(userID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": userID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		log.Println(err)
		return "", err
	}

	return tokenString, nil
}

// VerifyToken verifies the JWT token
func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Println(err)
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", errors.New("invalid token")
	}

	userID, ok := claims["username"].(string)
	if !ok {
		return "", errors.New("invalid userName in token")
	}

	return userID, nil
}
