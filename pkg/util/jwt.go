package util

import (
	"fmt"
	"gogin/example/pkg/setting"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtSecret = []byte(setting.JwtSecret)



func GenerateToken(username string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username" : username,
		"expired_at" : time.Now().Add(3 * time.Hour),
	})

	tokenString, err := token.SignedString(JwtSecret)

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return tokenString
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return JwtSecret, nil
	})

	if err != nil {
		return nil, err
	}
	
	claims := token.Claims.(jwt.MapClaims)

	return claims, nil
}