package routers

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/atharvYadavXperate/xlms/backend/handlers"
	db "github.com/atharvYadavXperate/xlms/database"
	customerror "github.com/atharvYadavXperate/xlms/database/errors"
	cutomeerror "github.com/atharvYadavXperate/xlms/database/errors"
)

func SetManger(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	manager, err := strconv.Atoi(query.Get("manager"))
	if err != nil {
		appErr := customerror.ErrBadRequest("Manager id filed is required", err)
		customerror.HandleError(w, appErr)
		return
	}
	user, err := strconv.Atoi(query.Get("user"))
	if err != nil {
		appErr := customerror.ErrBadRequest("User id filed is required", err)
		customerror.HandleError(w, appErr)
		return
	}
	log.Println("Manager: ", manager, " User: ", user)
	err = db.SetManager(context.Background(), int64(manager), int64(user))
	if err != nil {
		log.Println(err)
		appErr := cutomeerror.ErrBadRequest("Failed to set manager please check manager", err)
		cutomeerror.HandleError(w, appErr)
		return
	}
	handlers.Response(w, http.StatusAccepted, "Assigned manager to user", nil)
}
