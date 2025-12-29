package models

import "time"

type LeaveType int
type Role int64
type Time int

const (
	MORNING     Time = 1
	AFTERNOON   Time = 2
	END_OF_DAY  Time = 3
	AFTER_LUNCH Time = 4
)

const (
	AdminRole   Role = 1
	ManagerRole Role = 2
	UserRole    Role = 3
)

const (
	ExamLeave       LeaveType = 1
	Holiday         LeaveType = 2
	OptionalHoliday LeaveType = 3
	UnpaidLeave     LeaveType = 4
	WorkFromHome    LeaveType = 5
	UKShift         LeaveType = 6
)

type User struct {
	ID            int64     `json:"id" db:"id"`
	FullName      string    `json:"full_name" db:"full_name"`
	Email         string    `json:"email" db:"email"`
	IsApproved    bool      `json:"is_approved" db:"is_approved"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	ReportingToID *int64    `json:"reporting_to,omitempty" db:"reporting_to"`
	RoleId        Role      `json:"role_id" db:"role_id"`
	Role          string    `json:"role,omitempty" db:"role,omitempty"`
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
	StartAt    Time      `json:"start_at" db:"start_id"`
	EndAt      Time      `json:"end_at" db:"end_id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
}
