package domain

import (
	"time"
)

type TokenInfo struct {
	Exp         time.Time     `json:"exp"`
	Remaining   time.Duration `json:"remaining"`
	IsRefreshed bool          `json:"is_refreshed"`
	Role        string        `json:"role"`
	UserID      int64         `json:"user_id"`
}

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type TokenProvider interface {
	CreateTokenPair(userID int64, role string) (*TokenPair, error)
	CreateToken(userID int64, role string, accessToken bool) (string, error)
	ParseToken(token string) (*CustomClaims, error)
}
