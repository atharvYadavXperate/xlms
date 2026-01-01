package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/atharvYadavXperate/xlms/backend/handlers"
	"github.com/atharvYadavXperate/xlms/backend/utils"
	db "github.com/atharvYadavXperate/xlms/database"
	cutomeerror "github.com/atharvYadavXperate/xlms/database/errors"
	models "github.com/atharvYadavXperate/xlms/types"
)

func ApplyForLeave(w http.ResponseWriter, r *http.Request) {
	var leave models.Leave
	err := json.NewDecoder(r.Body).Decode(&leave)
	if err != nil {
		appErr := cutomeerror.ErrBadRequest("Invalid json formate", cutomeerror.InvalidJson)
		cutomeerror.HandleError(w, appErr)
		return
	}
	if leave.UserID == 0 || leave.FromDate.IsZero() || leave.ToDate.IsZero() || leave.LeaveType == 0 || leave.StartAt == 0 || leave.EndAt == 0 {
		appErr := cutomeerror.ErrBadRequest("All fields are required", cutomeerror.FieldsAreRequired)
		cutomeerror.HandleError(w, appErr)
		return
	}
	var appleLeave models.Leave
	appleLeave.UserID = leave.UserID
	appleLeave.FromDate = leave.FromDate
	appleLeave.ToDate = leave.ToDate
	appleLeave.LeaveType = leave.LeaveType
	appleLeave.StartAt = leave.StartAt
	appleLeave.EndAt = leave.EndAt
	appliedLeave, err := db.CreateNewLeave(context.Background(), appleLeave)
	if err != nil {
		appErr := cutomeerror.ErrInternal(err)
		cutomeerror.HandleError(w, appErr)
		return
	}
	handlers.Response(w, http.StatusAccepted, "Leave applied", appliedLeave)
}

func ApproveLeave(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	leaveId, err := strconv.Atoi(query.Get("leave_id"))
	if err != nil {
		appErr := cutomeerror.ErrBadRequest("Failed to get leave id", cutomeerror.FieldsAreRequired)
		cutomeerror.HandleError(w, appErr)
		return
	}
	var userInfo models.AccessToken
	userInfo, err = utils.GetUserValuesFromAccessToken(r)
	if err != nil {
		appErr := cutomeerror.ErrForbidden(err)
		cutomeerror.HandleError(w, appErr)
		return
	}
	approved, err := db.LeaveApprove(context.Background(), int64(leaveId), int64(userInfo.UserId))
	if err != nil {
		appErr := cutomeerror.ErrInternal(err)
		cutomeerror.HandleError(w, appErr)
		return
	}
	handlers.Response(w, http.StatusAccepted, "Leave approved", approved)
}
