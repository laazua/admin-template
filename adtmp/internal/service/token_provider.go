package service

import (
	"errors"
	"time"

	"adtmp/internal/domain/entities"
	"adtmp/pkg/config"

	"github.com/golang-jwt/jwt/v5"
)

type tokenProvider struct {
	secretKey []byte
	expiresIn time.Duration
}

func NewJWTProvider() TokenProvider {
	return &tokenProvider{
		secretKey: []byte(config.Get().SecretKey),
		expiresIn: config.Get().ExpiredTime,
	}
}

func (tokenProvider *tokenProvider) GenerateToken(user *entities.User) (string, error) {
	claims := &entities.TokenClaims{
		Email:    user.Email,
		Username: user.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenProvider.expiresIn)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(tokenProvider.secretKey)
}

func (tokenProvider *tokenProvider) ValidateToken(tokenString string) (*entities.TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &entities.TokenClaims{}, func(token *jwt.Token) (any, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return tokenProvider.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*entities.TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
