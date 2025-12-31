package middleware

import (
	"context"
	"net/http"

	"github.com/atharvYadavXperate/xlms/backend/utils"
	cutomeerror "github.com/atharvYadavXperate/xlms/database/errors"
)

func AuthMiddleWare(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		accessToken, err := r.Cookie("access_token")
		if err != nil {
			appErr := cutomeerror.ErrForbidden(cutomeerror.UnauthorizedAccess)
			cutomeerror.HandleError(w, appErr)
			return
		}

		claims, err := utils.VerifyAccessToken(accessToken.Value)
		if err != nil {
			appErr := cutomeerror.ErrForbidden(cutomeerror.UnauthorizedAccess)
			cutomeerror.HandleError(w, appErr)
			return
		}

		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
