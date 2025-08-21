package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	UserID       uint   `json:"uid"`
	Email        string `json:"email"`
	Role         string `json:"role"`
	TokenVersion int    `json:"tv"`
	jwt.RegisteredClaims
}

func secret() []byte {
	sec := os.Getenv("JWT_SECRET")
	if sec == "" {
		sec = "DEV_ONLY_CHANGE_ME" // dev fallback; ganti di prod!
	}
	return []byte(sec)
}

func GenerateToken(userID uint, email, role string, tokenVersion int, ttl time.Duration) (string, time.Time, error) {
	exp := time.Now().Add(ttl)
	claims := CustomClaims{
		UserID:       userID,
		Email:        email,
		Role:         role,
		TokenVersion: tokenVersion,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := tok.SignedString(secret())
	return s, exp, err
}

func ParseToken(tokenStr string) (*CustomClaims, error) {
	tok, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return secret(), nil
	})
	if err != nil {
		return nil, err
	}
	if !tok.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}
	cc, _ := tok.Claims.(*CustomClaims)
	return cc, nil
}
