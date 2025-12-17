package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

var JWT_SECRET = []byte("thisisverysecurekey")

type CustomClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}
