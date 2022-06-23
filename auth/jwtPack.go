package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJWT(email, password string) (string, error) {
	var JwtKey = []byte("my_secret_key")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["email"] = email
	claims["password"] = password
	claims["expired"] = time.Now().Add(time.Hour * 12).Unix()

	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
