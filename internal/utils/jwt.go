package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string = os.Getenv("JWT_SECRET")

func GenerateJWT(userID int, username string, role string) (string, error) {
	if jwtSecret == "" {
		return "",fmt.Errorf("JWT_SECRET IS EMPTY")
	}
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"role": role,
		
		"exp":      time.Now().Add(time.Hour * 24).Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
