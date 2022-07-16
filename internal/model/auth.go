package model

import "github.com/dgrijalva/jwt-go"

type Token struct {
	ID        string
	SessionId string
	JWTToken  jwt.StandardClaims
}
