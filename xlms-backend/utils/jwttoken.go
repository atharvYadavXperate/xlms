package utils

import (
	"fmt"
	"time"

	"github.com/atharvYadavXperate/xlms/backend/handlers"
	models "github.com/atharvYadavXperate/xlms/types"
	"github.com/golang-jwt/jwt/v5"
)

type RefreshClaims struct {
	UserID int64
	Email  string
	RoleId int
}

func CreateAccessToken(userID int64, email string, roleId int) (string, error) {
	secret := handlers.AccessTokenSecret
	if secret == "" {
		return "", fmt.Errorf("ACCESS_TOKEN_SECRET not set")
	}

	claims := models.AccessToken{
		Email:  email,
		RoleId: roleId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(3 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secret))
}

func VerifyAccessToken(tokenString string) (*models.AccessToken, error) {
	secret := handlers.AccessTokenSecret

	claims := &models.AccessToken{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method")
	}

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid access token")
	}

	return claims, nil
}

func GenerateRefreshToken(userID int64, email string, roleId int) (string, error) {
	secret := handlers.RefreshTokenSecret
	if secret == "" {
		return "", fmt.Errorf("REFRESH_TOKEN_SECRET not set")
	}

	claims := jwt.MapClaims{
		"sub":   userID,
		"email": email,
		"role":  roleId,
		"type":  "refresh",
		"exp":   time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ValidateRefreshToken(tokenStr string) (int64, error) {
	secret := handlers.RefreshTokenSecret
	if secret == "" {
		return 0, fmt.Errorf("REFRESH_TOKEN_SECRET not set")
	}

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return 0, fmt.Errorf("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, fmt.Errorf("invalid claims format")
	}

	if claims["type"] != "refresh" {
		return 0, fmt.Errorf("not a refresh token")
	}

	sub, ok := claims["sub"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid sub claim")
	}

	return int64(sub), nil
}
