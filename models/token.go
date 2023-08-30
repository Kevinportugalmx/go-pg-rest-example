package models

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	Active bool   `json:"active"`
	Scope  string `json:"scope"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

type ResponseToken struct {
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expiresAt"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token"`
}
