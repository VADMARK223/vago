package token

import (
	"errors"
	"fmt"
	"strconv"
	"time"
	"vago/internal/config/code"
	"vago/internal/config/config"
	"vago/internal/domain"

	"github.com/golang-jwt/jwt/v5"
)

type JWTProvider struct {
	Port       string
	secret     string
	accessTTL  time.Duration
	refreshTTL time.Duration
}

func NewJWTProvider(cfg *config.Config) *JWTProvider {
	return &JWTProvider{
		secret:     cfg.JwtSecret,
		accessTTL:  cfg.AccessTTLDuration(),
		refreshTTL: cfg.RefreshTTLDuration(),
		Port:       cfg.Port,
	}
}

func (j *JWTProvider) CreateTokenPair(userID int64, role string, username string) (*domain.TokenPair, error) {
	access, err := j.CreateToken(userID, role, username, true)
	if err != nil {
		return nil, err
	}

	refresh, err := j.CreateToken(userID, role, username, false)
	if err != nil {
		return nil, err
	}

	return &domain.TokenPair{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}

func (j *JWTProvider) CreateToken(userID int64, role string, username string, accessToken bool) (string, error) {
	now := time.Now()

	var duration time.Duration
	if accessToken {
		duration = j.accessTTL
	} else {
		duration = j.refreshTTL
	}

	claims := domain.CustomClaims{
		Role:     role,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(int(userID)),
			ExpiresAt: jwt.NewNumericDate(now.Add(duration)),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now), // Делает токен “активным не раньше определённого момента”
			Issuer:    code.Vago,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secret))
}

func (j *JWTProvider) ParseToken(tokenStr string) (*domain.CustomClaims, error) {
	if tokenStr == "" {
		return nil, errors.New("token is empty")
	}

	token, err := jwt.ParseWithClaims(tokenStr, &domain.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("parse error: %w", err)
	}

	claims, ok := token.Claims.(*domain.CustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
