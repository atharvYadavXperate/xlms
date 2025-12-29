package database

import (
	"context"
	"fmt"
)

func SeedOperation(ctx context.Context) {
	fmt.Println("Performing seed Operations")
	if err := seedRoles(ctx); err != nil {
		panic(err)
	}
	if err := seedLeaveTypes(ctx); err != nil {
		panic(err)
	}
	if err := timeSeeding(ctx); err != nil {
		panic(err)
	}
	fmt.Println("Seed Operation Completed...")
}

func seedRoles(ctx context.Context) error {
	fmt.Println("Adding Roles")
	query := `
		INSERT INTO roles (id, role)
		VALUES
			(1, 'ADMIN'),
			(2, 'MANAGER'),
			(3, 'USER')
		ON CONFLICT (id) DO NOTHING
	`
	_, err := Pool.Exec(ctx, query)
	return err
}

func seedLeaveTypes(ctx context.Context) error {
	fmt.Println("Adding Leave Types")

	query := `
		INSERT INTO leave_type (leave_id, leave_name) 
		VALUES
			(1, 'Exam Leave'),
			(2, 'Holiday'),
			(3, 'Optional Holiday'),
			(4, 'Unpaid Leave'),
			(5, 'Work From Home'),
			(6, 'Work in Shift (UK)')
		ON CONFLICT (leave_id) DO NOTHING
	`

	_, err := Pool.Exec(ctx, query)
	return err
}

func timeSeeding(ctx context.Context) error {
	fmt.Println("Adding Start time")
	query := `
		INSERT INTO start_at(start_id, time)
		VALUES
			(1, 'Morning'),
			(2, 'Afternoon')
		ON CONFLICT (start_id) DO NOTHING
	`
	_, err := Pool.Exec(ctx, query)

	if err != nil {
		return err
	}
	fmt.Println("Adding End time")
	query = `
		INSERT INTO end_at(end_id, time)
		VALUES
			(3, 'Lunch Time'),
			(4, 'End of Day')
		ON CONFLICT (end_id) DO NOTHING
	`
	_, err = Pool.Exec(ctx, query)
	return err
}
