package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/atharvYadavXperate/xlms/backend/utils"
	cutomeerror "github.com/atharvYadavXperate/xlms/database/errors"
)

func MangerAndAdminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Auth Middleware")
		accessToken, err := r.Cookie("access_token")
		if err != nil {
			appErr := cutomeerror.ErrForbidden(cutomeerror.UnauthorizedAccess)
			cutomeerror.HandleError(w, appErr)
			return
		}

		claims, err := utils.VerifyAccessToken(accessToken.Value)
		log.Println(claims)
		if err != nil {
			appErr := cutomeerror.ErrForbidden(cutomeerror.UnauthorizedAccess)
			cutomeerror.HandleError(w, appErr)
			return
		}

		if claims.RoleId == 1 {
			appErr := cutomeerror.ErrForbidden(cutomeerror.UnauthorizedAccess)
			cutomeerror.HandleError(w, appErr)
			return
		}
		ctx := context.WithValue(r.Context(), "user", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
