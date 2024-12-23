// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: driver.sql

package queries

import (
	"context"
)

const createDriver = `-- name: CreateDriver :one
INSERT INTO drivers (user_id, vehicle, availability_status)
VALUES (?1, ?2, ?3)
RETURNING id
`

type CreateDriverParams struct {
	UserID             int64
	Vehicle            string
	AvailabilityStatus string
}

func (q *Queries) CreateDriver(ctx context.Context, arg CreateDriverParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createDriver, arg.UserID, arg.Vehicle, arg.AvailabilityStatus)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateDriver = `-- name: UpdateDriver :one
UPDATE drivers
SET vehicle = ?1,
    availability_status = ?2
WHERE id = ?3
RETURNING id
`

type UpdateDriverParams struct {
	Vehicle            string
	AvailabilityStatus string
	ID                 int64
}

func (q *Queries) UpdateDriver(ctx context.Context, arg UpdateDriverParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateDriver, arg.Vehicle, arg.AvailabilityStatus, arg.ID)
	var id int64
	err := row.Scan(&id)
	return id, err
}
