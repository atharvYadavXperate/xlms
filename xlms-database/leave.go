package database

import (
	"context"
	"errors"

	models "github.com/atharvYadavXperate/xlms/types"
)

func CreateNewLeave(ctx context.Context, leave models.Leave) (models.Leave, error) {
	query := `
		INSERT INTO leaves (
			user_id,
			from_date,
			to_date,
			leave_type,
			reason,
			start_time,
			end_time
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, user_id, from_date, to_date, leave_type, reason, start_time, end_time, is_approved, created_at;
	`

	var newLeave models.Leave
	err := Pool.QueryRow(
		ctx,
		query,
		leave.UserID,
		leave.FromDate,
		leave.ToDate,
		leave.LeaveType,
		leave.Reason,
		leave.StartAt,
		leave.EndAt,
	).Scan(
		&newLeave.ID,
		&newLeave.UserID,
		&newLeave.FromDate,
		&newLeave.ToDate,
		&newLeave.LeaveType,
		&newLeave.Reason,
		&newLeave.StartAt,
		&newLeave.EndAt,
		&newLeave.IsApproved,
		&newLeave.CreatedAt,
	)

	return newLeave, err
}

func LeaveApprove(ctx context.Context, leaveId int64, managerId int64) (bool, error) {
	query := `
		UPDATE leaves l
		SET is_approved = true
		FROM users u
		WHERE l.user_id = u.id
		  AND u.reporting_to = $1
		  AND l.id = $2
		RETURNING l.id;
	`

	var updatedId int64
	err := Pool.QueryRow(ctx, query, managerId, leaveId).Scan(&updatedId)

	if err != nil {
		return false, err
	}

	return true, nil
}

func DeleteLeave(ctx context.Context, leaveId, requesterId int64) error {
	query := `
		DELETE FROM leaves l
		USING users u
		WHERE l.user_id = u.id
		  AND l.id = $1
		  AND (u.id = $2 OR u.reporting_to = $2);
	`

	tag, err := Pool.Exec(ctx, query, leaveId, requesterId)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return errors.New("not authorized to delete this leave")
	}

	return nil
}
