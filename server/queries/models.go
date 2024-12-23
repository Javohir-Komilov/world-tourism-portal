// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package queries

import (
	"database/sql"
	"time"
)

type Driver struct {
	ID                 int64
	UserID             int64
	Vehicle            string
	AvailabilityStatus string
	CreatedAt          sql.NullTime
}

type Hotel struct {
	ID        int64
	ManagerID int64
	Name      string
	Location  string
	Rating    sql.NullFloat64
	CreatedAt sql.NullTime
}

type HotelBooking struct {
	ID        int64
	HotelID   int64
	TouristID int64
	CheckIn   time.Time
	CheckOut  time.Time
	CreatedAt sql.NullTime
}

type TourPackage struct {
	ID          int64
	AgentID     int64
	Name        string
	Description sql.NullString
	Price       float64
	StartDate   time.Time
	EndDate     time.Time
	CreatedAt   sql.NullTime
}

type TripOrder struct {
	ID          int64
	DriverID    int64
	TouristID   int64
	Destination string
	Status      string
	CreatedAt   sql.NullTime
}

type User struct {
	ID           int64
	Username     string
	PasswordHash string
	Role         string
	CreatedAt    sql.NullTime
}
