package database

import (
	"context"
	"errors"
	"fmt"

	u "github.com/atharvYadavXperate/xlms/types"
)

func CreateUser(ctx context.Context, fullName, email string, role_id int) (int64, error) {
	if Pool == nil {
		return 0, errors.New("database pool is not initialized")
	}

	var userID int64
	query := `
		INSERT INTO users (full_name, email, role_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := Pool.QueryRow(ctx, query, fullName, email, role_id).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return userID, nil
}

func GetUserByID(ctx context.Context, id int64) (u.User, error) {
	var user u.User
	query := `
		SELECT 
			a.id,
			a.full_name,
			a.email,
			a.is_approved,
			a.created_at,
			a.reporting_to,
			r.id   AS role_id,
			r.role AS role
		FROM users a
		LEFT JOIN roles r ON a.role_id = r.id
		WHERE a.id = $1;
	`

	err := Pool.QueryRow(ctx, query, id).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.IsApproved,
		&user.CreatedAt,
		&user.ReportingToID,
		&user.RoleId,
		&user.Role,
	)

	if err != nil {
		return u.User{}, err
	}

	return user, nil
}

func GetUserByEmail(ctx context.Context, email string) (u.User, error) {
	var user u.User

	query := `
		SELECT 
			a.id,
			a.full_name,
			a.email,
			a.is_approved,
			a.created_at,
			a.reporting_to,
			r.id   AS role_id,
			r.role AS role
		FROM users a
		LEFT JOIN roles r ON a.role_id = r.id
		WHERE a.email = $1;
	`

	err := Pool.QueryRow(ctx, query, email).Scan(
		&user.ID,
		&user.FullName,
		&user.Email,
		&user.IsApproved,
		&user.CreatedAt,
		&user.ReportingToID,
		&user.RoleId,
		&user.Role,
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
			&user.ReportingToID,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUsers(ctx context.Context, page int, limit int) ([]u.User, error) {
	const maxlimit = 500
	if limit > maxlimit {
		limit = maxlimit
	}
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}
	query := `
		SELECT id, full_name, email, role_id, is_approved, created_at, reporting_to
		FROM users
		ORDER BY id
		LIMIT $1 OFFSET $2
	`

	rows, err := Pool.Query(ctx, query, limit, offset)
	defer rows.Close()
	if err != nil {
		return nil, err
	}
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
			&user.ReportingToID,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func UpdateUser(ctx context.Context, id int64, fullName string, email string, role int) (u.User, error) {
	var user u.User
	query := `
		UPDATE users
		SET full_name = $1,
		    email = $2,
		    role_id = $3
		WHERE id = $4 
		RETURNING id, full_name, email, role_id, is_approved, created_at
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
	fmt.Println(err)
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
		WHERE id = $1 
		RETURNING id, full_name, email, role_id, is_approved, created_at
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

func SoftDeleteUser(ctx context.Context, id int64) (u.User, error) {
	var user u.User
	query := `
		UPDATE users
		SET is_approved = false
		WHERE id = $1 
		RETURNING id, full_name, email, role_id, is_approved, created_at
	`
	err := Pool.QueryRow(ctx, query, id).Scan(&user.ID, &user.FullName, &user.Email, &user.RoleId, &user.IsApproved, &user.CreatedAt)
	if err != nil {
		return u.User{}, err
	}
	return user, nil
}
