package jwtHelper

import "github.com/dgrijalva/jwt-go"

func GenerateToken(claims *jwt.Token, secret string) string{
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token , _ := claims.SignedString(hmacSecret)
	return token
}