package database

import (
	"context"
	"errors"

	u "github.com/atharvYadavXperate/xlms/types"
)

func CreateUser(ctx context.Context, fullName, email, role string) (int64, error) {
	var userID int64

	query := `
		INSERT INTO users (full_name, email, role)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := Pool.QueryRow(ctx, query, fullName, email, role).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func GetUserByID(ctx context.Context, id int64) (u.User, error) {
	var user u.User
	query := `
		SELECT id, full_name, email, role, is_approved, created_at, updated_at
		FROM users
		WHERE id = $1 AND deleted_at IS NULL
	`
	err := Pool.QueryRow(ctx, query, id).
		Scan(
			&user.ID,
			&user.FullName,
			&user.Email,
			&user.RoleId,
			&user.IsApproved,
			&user.CreatedAt,
		)
	if err != nil {
		return u.User{}, err
	}
	return user, nil
}

func GetUsersByManager(ctx context.Context, id int64, page int) ([]u.User, error) {
	const limit = 10
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}
	query := `
		SELECT id, full_name, email, role_id, is_approved, created_at, reporting_to
		FROM users
		WHERE reporting_to = $1
		ORDER BY id
		LIMIT $2 OFFSET $3
	`
	rows, err := Pool.Query(ctx, query, id, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []u.User{}
	for rows.Next() {
		var user u.User
		err := rows.Scan(
			&user.ID,
			&user.FullName,
			&user.Email,
			&user.RoleId,
			&user.IsApproved,
			&user.CreatedAt,
			&user.ReportingTo,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func UpdateUser(ctx context.Context, id int64, fullName string, email string, role string) (u.User, error) {
	var user u.User
	query := `
		UPDATE users
		SET full_name = $1,
		    email = $2,
		    role = $3
		WHERE id = $4 AND deleted_at IS NULL
		RETURNING id, full_name, email, role, is_approved, created_at, updated_at
	`
	err := Pool.QueryRow(ctx, query, fullName, email, role, id).
		Scan(
			&user.ID,
			&user.FullName,
			&user.Email,
			&user.RoleId,
			&user.IsApproved,
			&user.CreatedAt,
		)

	if err != nil {
		return u.User{}, err
	}
	return user, nil
}

func VerifyUser(ctx context.Context, id int64) (u.User, error) {
	var user u.User
	query := `
		UPDATE users
		SET is_approved = true
		WHERE id = $1 AND deleted_at IS NULL
		RETURNING id, full_name, email, role, is_approved, created_at, updated_at
	`
	err := Pool.QueryRow(ctx, query, id).
		Scan(
			&user.ID,
			&user.FullName,
			&user.Email,
			&user.RoleId,
			&user.IsApproved,
			&user.CreatedAt,
		)
	if err != nil {
		return u.User{}, err
	}
	return user, nil
}

func SoftDeleteUser(ctx context.Context, id int64) error {
	query := `
		UPDATE users
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	cmd, err := Pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}
	if cmd.RowsAffected() == 0 {
		return errors.New("user not found or already deleted")
	}
	return nil
}
