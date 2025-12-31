package routers

import (
	"context"
	"net/http"
	"time"

	"github.com/atharvYadavXperate/xlms/backend/handlers"
	"github.com/atharvYadavXperate/xlms/backend/utils"
	db "github.com/atharvYadavXperate/xlms/database"
	cutomeerror "github.com/atharvYadavXperate/xlms/database/errors"
)

func RefreshAccessToken(w http.ResponseWriter, r *http.Request) {
	refreshCookie, err := r.Cookie("refresh_token")
	if err != nil {
		cutomeerror.HandleError(w, cutomeerror.ErrForbidden(cutomeerror.UnauthorizedAccess))
		return
	}

	userID, err := utils.ValidateRefreshToken(refreshCookie.Value)
	if err != nil {
		cutomeerror.HandleError(w, cutomeerror.ErrForbidden(cutomeerror.UnauthorizedAccess))
		return
	}

	user, err := db.GetUserByID(context.Background(), userID)
	if err != nil {
		cutomeerror.HandleError(w, cutomeerror.ErrForbidden(cutomeerror.UnauthorizedAccess))
		return
	}

	newAccessToken, err := utils.CreateAccessToken(user.ID, user.Email, int(user.RoleId))
	if err != nil {
		cutomeerror.HandleError(w, cutomeerror.ErrInternal(err))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "access_token",
		Value:    newAccessToken,
		Path:     "/",
		Expires:  time.Now().Add(3 * time.Minute),
		HttpOnly: true,
		Secure:   false, // localhost
		SameSite: http.SameSiteLaxMode,
	})

	handlers.Response(w, http.StatusOK, "Access token refreshed", nil)
}
