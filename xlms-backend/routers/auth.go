package routers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/atharvYadavXperate/xlms/backend/handlers"
	"github.com/atharvYadavXperate/xlms/backend/utils"
	d "github.com/atharvYadavXperate/xlms/database"
	cutomeerror "github.com/atharvYadavXperate/xlms/database/errors"
	u "github.com/atharvYadavXperate/xlms/types"
	"github.com/jackc/pgx/v5/pgconn"
)

var OTPs = make(map[string]int)

func Register(w http.ResponseWriter, r *http.Request) {

	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		err := cutomeerror.ErrBadRequest(
			"Content-Type must be application/json",
			cutomeerror.ContentTypeConflictMustJson,
		)
		cutomeerror.HandleError(w, err)
		return
	}

	var user u.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		err := cutomeerror.ErrBadRequest(
			"Invalid JSON format",
			cutomeerror.InvalidJson,
		)
		cutomeerror.HandleError(w, err)
		return
	}

	if user.FullName == "" || user.Email == "" || user.RoleId == 0 {
		err := cutomeerror.ErrBadRequest(
			"Required fields are missing",
			cutomeerror.FieldsAreRequired,
		)
		cutomeerror.HandleError(w, err)
		return
	}

	id, err := d.CreateUser(
		context.Background(),
		user.FullName,
		user.Email,
		int(user.RoleId),
	)
	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			err := cutomeerror.ErrDuplication(errors.New("email already exists"))
			cutomeerror.HandleError(w, err)
			return
		}
		cutomeerror.HandleError(w, err)
		return
	}

	user.ID = id
	handlers.Response(w, http.StatusCreated, "User registered successfully", user)
}

func GenerateOTP(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		err := cutomeerror.ErrBadRequest(
			"Content-Type must be application/json",
			cutomeerror.ContentTypeConflictMustJson,
		)
		cutomeerror.HandleError(w, err)
		return
	}

	var req struct {
		Email string `json:"email"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := cutomeerror.ErrBadRequest("Invalid JSON format", cutomeerror.InvalidJson)
		cutomeerror.HandleError(w, err)
		return
	}

	if req.Email == "" {
		err := cutomeerror.ErrBadRequest("Email is required", cutomeerror.FieldsAreRequired)
		cutomeerror.HandleError(w, err)
		return
	}

	user, err := d.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		err := cutomeerror.ErrNotFound("User not found", err)
		cutomeerror.HandleError(w, err)
		return
	}

	if !user.IsApproved {
		err := cutomeerror.ErrForbidden(cutomeerror.UserNotApproved)
		cutomeerror.HandleError(w, err)
		return
	}

	otp := utils.Random4Digit()
	log.Println("Generated OTP:", otp) // TODO: send email

	OTPs[req.Email] = otp

	handlers.Response(w, http.StatusOK, "OTP sent successfully", nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if !strings.Contains(r.Header.Get("Content-Type"), "application/json") {
		err := cutomeerror.ErrBadRequest(
			"Content-Type must be application/json",
			cutomeerror.ContentTypeConflictMustJson,
		)
		cutomeerror.HandleError(w, err)
		return
	}

	var req struct {
		Email string `json:"email"`
		Otp   int    `json:"otp"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		err := cutomeerror.ErrBadRequest("Invalid JSON format", cutomeerror.InvalidJson)
		cutomeerror.HandleError(w, err)
		return
	}
	log.Println("Email: ", req.Email)
	log.Println("otp: ", req.Otp)
	if req.Email == "" || req.Otp < 1000 || req.Otp > 9999 {
		err := cutomeerror.ErrBadRequest("Invalid email or OTP", cutomeerror.FieldsAreRequired)
		cutomeerror.HandleError(w, err)
		return
	}

	user, err := d.GetUserByEmail(context.Background(), req.Email)
	if err != nil {
		err := cutomeerror.ErrNotFound("User not found", err)
		cutomeerror.HandleError(w, err)
		return
	}

	if !user.IsApproved {
		err := cutomeerror.ErrForbidden(cutomeerror.UserNotApproved)
		cutomeerror.HandleError(w, err)
		return
	}

	storedOtp, ok := OTPs[req.Email]
	if !ok || storedOtp != req.Otp {
		err := cutomeerror.ErrForbidden(cutomeerror.InvalidOtp)
		cutomeerror.HandleError(w, err)
		return
	}

	loginUser := u.UserAuthRes{
		Email: user.Email,
		Role:  int(user.RoleId),
		Token: "mytoken",
	}

	handlers.Response(w, http.StatusOK, "Login successful", loginUser)
}
