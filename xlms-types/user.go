package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserAuthRes struct {
	Email string `json:"email"`
	Role  int    `json:"role_id"`
	Token string `json:"token"`
}

type RefreshToken struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Token     string    `db:"token"`
	ExpiresAt time.Time `db:"expires_at"`
	Revoked   bool      `db:"revoked"`
	CreatedAt time.Time `db:"created_at"`
}

type AccessToken struct {
	Email  string `json:"email"`
	RoleId int    `json:"role_id"`
	jwt.RegisteredClaims
}
