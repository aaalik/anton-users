package jwtgen

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"
)

func (j *JwtUtils) GenerateToken(claims jwt.MapClaims, secretKey string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}

func (j *JwtUtils) ParseToken(tokenString, secretKey string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
}
