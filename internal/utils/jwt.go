package utils

import (
	"time"
	"os"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret string = os.Getenv("JWT_SECRET")

func generateJWT(userID int, username string) (string, error) {
	// Creamos claims (datos que queremos guardar en el token)
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // expira en 24h
	}

	// Creamos el token con el método de firma HMAC SHA256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmamos el token con la clave secreta
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
