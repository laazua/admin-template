package entities

import "github.com/golang-jwt/jwt/v5"

type TokenClaims struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	jwt.RegisteredClaims
}
