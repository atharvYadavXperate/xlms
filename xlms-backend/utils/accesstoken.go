package utils

import (
	"net/http"

	models "github.com/atharvYadavXperate/xlms/types"
)

func GetUserValuesFromAccessToken(r *http.Request) (models.AccessToken, error) {
	accessToken, err := r.Cookie("access_token")
	if err != nil {
		return models.AccessToken{}, err
	}
	claims, err := VerifyAccessToken(accessToken.Value)
	if err != nil {
		return models.AccessToken{}, err
	}
	var aToken models.AccessToken
	aToken.UserId = claims.UserId
	aToken.Email = claims.Email
	aToken.RoleId = claims.RoleId
	return aToken, nil
}
