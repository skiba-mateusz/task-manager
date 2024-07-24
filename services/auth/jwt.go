package auth

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/skiba-mateusz/task-manager/config"
)

func GenerateJWT(secret []byte, userID int) (string, error) {
	duration, err := time.ParseDuration(config.Envs.JWTExpiration) 
	if err != nil {
		return "", fmt.Errorf("error parsing time: %v", err)
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": 	 strconv.Itoa(userID),
		"expiresAt": time.Now().Add(duration).Unix(),
	})

	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}