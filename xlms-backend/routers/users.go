package routers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/atharvYadavXperate/xlms/backend/handlers"
	db "github.com/atharvYadavXperate/xlms/database"
	cutomeerror "github.com/atharvYadavXperate/xlms/database/errors"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	page, err := strconv.Atoi(query.Get("page"))
	if err != nil {
		page = 1
	}
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil {
		limit = 10
	}
	users, err := db.GetUsers(context.Background(), page, limit)
	if err != nil {
		err := cutomeerror.ErrInternal(err)
		cutomeerror.HandleError(w, err)
		log.Println(err.Err)
		return
	}
	handlers.Response(w, http.StatusOK, "Users data", users)
}
