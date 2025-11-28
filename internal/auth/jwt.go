package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey     []byte
	tokenDuration time.Duration
}

type UserClaims struct {
	Username string `json:"username"`
	Role     string `json:"role,omitempty"`
	jwt.RegisteredClaims
}

func NewJWTManager(secret string, duration time.Duration) *JWTManager {
	return &JWTManager{
		secretKey:     []byte(secret),
		tokenDuration: duration,
	}
}

func (j *JWTManager) GenerateToken(username, role string) (string, error) {
	now := time.Now().UTC()
	claims := UserClaims{
		Username: username,
		Role:     role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(j.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			Subject:   username,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secretKey)
}

func (j *JWTManager) VerifyToken(accessToken string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return j.secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*UserClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, err
}
