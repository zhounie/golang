package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(userId string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := CustomClaims{
		UserID: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Subject:   userId,
			Issuer:    "auth-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	tokenString, err := token.SignedString(JWT_SECRET)
	if err != nil {
		return "", fmt.Errorf("生成 token 失败：%w", err)
	}
	return tokenString, nil
}

func ValidateToken(tokenString string) (*CustomClaims, error) {
	claims := &CustomClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("非预期的签名方法：%v", t.Header["alg"])
		}
		return JWT_SECRET, nil
	})

	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, fmt.Errorf("token 结构验证失败")
	}
	return claims, nil
}
