package utils

import (
	"errors"
	"go-backend-template/ent"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJwtToken(user *ent.User) (string, error) {
	claims := jwt.MapClaims{
		"user": user,
		"exp":  time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	log.Println("Secret Key", os.Getenv("JWT_SECRET"))

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func VerfyJwtToken(tokenString string) (jwt.MapClaims, error) {
	parseToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("jwt token expired")
		}
		return nil, errors.New("invalid jwt token")
	}

	if claims, ok := parseToken.Claims.(jwt.MapClaims); ok && parseToken.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid jwt token")
}
