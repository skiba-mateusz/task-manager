package auth

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/skiba-mateusz/task-manager/config"
	"github.com/skiba-mateusz/task-manager/models"
	"github.com/skiba-mateusz/task-manager/utils"
)

type userContextKey string

const (
	userKey userContextKey = "userID"
)

func AuthMiddleware(userStore models.UserStore) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenStr := r.Header.Get("Authorization")
			log.Println(tokenStr)

			token, err := validateJWT(tokenStr)
			if err != nil {
				log.Printf("failed to validate token: %v", err)
				permissionDenied(w)
				return
			}
	
			if !token.Valid {
				log.Println("invalid token")
				permissionDenied(w)
				return
			}
	
			claims := token.Claims.(jwt.MapClaims)
			str := claims["userID"].(string) 
	
			userID, err := strconv.Atoi(str)
			if err != nil {
				log.Printf("failed to convert userID to int")
				permissionDenied(w)
				return
			}

			user, err := userStore.GetUserByID(userID)
			if err != nil {
				log.Printf("failed to get user by id: %v", err)
				permissionDenied(w)
				return
			}

			ctx := r.Context()
			ctx = context.WithValue(ctx, userKey, user)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		})
	}
}

func GetUserFromContext(ctx context.Context) models.User {
	user, ok := ctx.Value(userKey).(models.User)
	if !ok {
		return models.User{}
	}
	return user
}

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

func validateJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(config.Envs.JWTSecret), nil	
	})
}

func permissionDenied(w http.ResponseWriter) {
	utils.WriteJSONError(w, http.StatusForbidden, fmt.Errorf("permision denied"))
} 