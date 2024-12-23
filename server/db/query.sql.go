// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
	"time"
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

const createHotel = `-- name: CreateHotel :one
INSERT INTO hotels (manager_id, name, location, rating)
VALUES (?1, ?2, ?3, ?4)
RETURNING id
`

type CreateHotelParams struct {
	ManagerID int64
	Name      string
	Location  string
	Rating    sql.NullFloat64
}

func (q *Queries) CreateHotel(ctx context.Context, arg CreateHotelParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createHotel,
		arg.ManagerID,
		arg.Name,
		arg.Location,
		arg.Rating,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createTourPackage = `-- name: CreateTourPackage :one
INSERT INTO tour_packages (agent_id, name, description, price, start_date, end_date)
VALUES (?1, ?2, ?3, ?4, ?5, ?6)
RETURNING id
`

type CreateTourPackageParams struct {
	AgentID     int64
	Name        string
	Description sql.NullString
	Price       float64
	StartDate   time.Time
	EndDate     time.Time
}

func (q *Queries) CreateTourPackage(ctx context.Context, arg CreateTourPackageParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createTourPackage,
		arg.AgentID,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.StartDate,
		arg.EndDate,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, password_hash, role)
VALUES (?1, ?2, ?3)
RETURNING id
`

type CreateUserParams struct {
	Username     string
	PasswordHash string
	Role         string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.PasswordHash, arg.Role)
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

const updateHotel = `-- name: UpdateHotel :one
UPDATE hotels
SET name = ?1,
    location = ?2,
    rating = ?3
WHERE id = ?4
RETURNING id
`

type UpdateHotelParams struct {
	Name     string
	Location string
	Rating   sql.NullFloat64
	ID       int64
}

func (q *Queries) UpdateHotel(ctx context.Context, arg UpdateHotelParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateHotel,
		arg.Name,
		arg.Location,
		arg.Rating,
		arg.ID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateTourPackage = `-- name: UpdateTourPackage :one
UPDATE tour_packages
SET name = ?1,
    description = ?2,
    price = ?3,
    start_date = ?4,
    end_date = ?5
WHERE id = ?6
RETURNING id
`

type UpdateTourPackageParams struct {
	Name        string
	Description sql.NullString
	Price       float64
	StartDate   time.Time
	EndDate     time.Time
	ID          int64
}

func (q *Queries) UpdateTourPackage(ctx context.Context, arg UpdateTourPackageParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateTourPackage,
		arg.Name,
		arg.Description,
		arg.Price,
		arg.StartDate,
		arg.EndDate,
		arg.ID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET username = ?1,
    password_hash = ?2,
    role = ?3
WHERE id = ?4
RETURNING id
`

type UpdateUserParams struct {
	Username     string
	PasswordHash string
	Role         string
	ID           int64
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (int64, error) {
	row := q.db.QueryRowContext(ctx, updateUser,
		arg.Username,
		arg.PasswordHash,
		arg.Role,
		arg.ID,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}
