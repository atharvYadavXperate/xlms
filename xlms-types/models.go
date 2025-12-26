package models

import "time"

type LeaveType string
type Role int64

const (
	ExamLeave       LeaveType = "Exam Leave"
	Holiday         LeaveType = "Holiday"
	OptionalHoliday LeaveType = "Optional Holiday"
	UnpaidLeave     LeaveType = "Unpaid Leave"
	WorkFromHome    LeaveType = "Work From Home"
	UKShift         LeaveType = "Working in Shift (UK)"
)

type User struct {
	ID          int64     `json:"id" db:"id"`
	FullName    string    `json:"full_name" db:"full_name"`
	Email       string    `json:"email" db:"email"`
	RoleId      Role      `json:"role_id" db:"role_id"`
	IsApproved  bool      `json:"is_approved" db:"is_approved"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	ReportingTo int64     `json:"reporting_to" db:reporting_to"`
}

type Team struct {
	ID        int64     `json:"id" db:"id"`
	TeamName  string    `json:"team_name" db:"team_name"`
	ManagerID int64     `json:"manager_id" db:"manager_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type Leave struct {
	ID         int64     `json:"id" db:"id"`
	UserID     int64     `json:"user_id" db:"user_id"`
	TeamID     int64     `json:"team_id" db:"team_id"`
	IsApproved bool      `json:"is_approved" db:"is_approved"`
	FromDate   time.Time `json:"from_date" db:"from_date"`
	ToDate     time.Time `json:"to_date" db:"to_date"`
	LeaveType  LeaveType `json:"leave_type" db:"leave_type"`
	Reason     string    `json:"reason" db:"reason"`
	StartAt    string    `json:"start_at" db:"start_at"`
	EndAt      string    `json:"end_at" db:"end_at"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
