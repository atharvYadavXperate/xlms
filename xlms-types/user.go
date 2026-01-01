package models

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserAuthRes struct {
	Email        string `json:"email"`
	Role         int    `json:"role_id"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type RefreshToken struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Token     string    `db:"token"`
	ExpiresAt time.Time `db:"expires_at"`
	Revoked   bool      `db:"revoked"`
	CreatedAt time.Time `db:"created_at"`
}
type RefreshClaims struct {
	UserID int64
	Email  string
	RoleId int
}

type AccessToken struct {
	UserId int    `json:"id"`
	Email  string `json:"email"`
	RoleId int    `json:"role_id"`
	jwt.RegisteredClaims
}
