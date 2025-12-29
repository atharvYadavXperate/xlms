package models

type UserAuthRes struct {
	Email string `json: "email"`
	Role  int    `json:"role_id"`
	Token string `json:"token"`
}
