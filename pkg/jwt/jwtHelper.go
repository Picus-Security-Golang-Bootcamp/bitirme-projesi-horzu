package jwtHelper

import (
	"encoding/json"

	"github.com/dgrijalva/jwt-go"
)

type DecodedToken struct {
	Iat    int    `json:"iat"`
	Role   string `json:"role"`
	UserId string `json:"userID"`
	Email  string `json:"email"`
	Iss    string `json:"iss"`
}

func GenerateToken(claims *jwt.Token, secret string) string {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)
	token, _ := claims.SignedString(hmacSecret)
	return token
}

func VerifyToken(token, secret string) *DecodedToken {
	hmacSecretString := secret
	hmacSecret := []byte(hmacSecretString)

	decoded, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return hmacSecret, nil
	})
	if err != nil {
		return nil
	}

	if !decoded.Valid {
		return nil
	}

	decodedClaims := decoded.Claims.(jwt.MapClaims)

	var decodedToken DecodedToken
	jsonString, _ := json.Marshal(decodedClaims)
	json.Unmarshal(jsonString, &decodedToken)

	return &decodedToken

}
