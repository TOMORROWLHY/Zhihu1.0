package model

import "github.com/dgrijalva/jwt-go"

type MyClaims struct {
	Username string
	jwt.StandardClaims
}
