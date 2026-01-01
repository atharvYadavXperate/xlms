package database

import "context"

func SetManager(ctx context.Context, managerId int64, userId int64) error {
	query := `
		UPDATE users u
			SET reporting_to = m.id
			FROM users m
			WHERE m.id = $1
  				AND m.role_id IN (1, 2)
 				AND u.id = $2;
	`
	_, err := Pool.Exec(ctx, query, managerId, userId)
	return err
}
