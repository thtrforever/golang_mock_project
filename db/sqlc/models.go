// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"time"
)

type BookingFlight struct {
	ID                   int64         `json:"id"`
	BookingTransactionID sql.NullInt64 `json:"booking_transaction_id"`
	Sequence             int32         `json:"sequence"`
	DepartureAirportCode string        `json:"departure_airport_code"`
	ArrivalAirportCode   string        `json:"arrival_airport_code"`
	DepartureDateTime    time.Time     `json:"departure_date_time"`
	ArrivalDateTime      time.Time     `json:"arrival_date_time"`
	BookingClass         string        `json:"booking_class"`
	CurrencyCode         string        `json:"currency_code"`
	Price                float64       `json:"price"`
}

type BookingPassenger struct {
	ID                   int64          `json:"id"`
	BookingTransactionID sql.NullInt64  `json:"booking_transaction_id"`
	Title                sql.NullString `json:"title"`
	LastName             sql.NullString `json:"last_name"`
	FirstName            sql.NullString `json:"first_name"`
	FfpNumber            sql.NullString `json:"ffp_number"`
	TicketNumber         sql.NullString `json:"ticket_number"`
}

type BookingTransaction struct {
	ID              int64          `json:"id"`
	ReservationCode sql.NullString `json:"reservation_code"`
	Status          sql.NullString `json:"status"`
	ErrorCode       sql.NullInt32  `json:"error_code"`
	ErrorMessage    sql.NullString `json:"error_message"`
	CreateDate      sql.NullTime   `json:"create_date"`
	ModifyDate      sql.NullTime   `json:"modify_date"`
}

type Customer struct {
	ID          int64          `json:"id"`
	Username    string         `json:"username"`
	Password    sql.NullString `json:"password"`
	FirstName   sql.NullString `json:"first_name"`
	LastName    sql.NullString `json:"last_name"`
	Address     sql.NullString `json:"address"`
	Email       sql.NullString `json:"email"`
	PhoneNumber sql.NullString `json:"phone_number"`
	Active      sql.NullBool   `json:"active"`
	CreateAt    sql.NullTime   `json:"create_at"`
	UpdateAt    sql.NullTime   `json:"update_at"`
}

type Flight struct {
	ID                     int64          `json:"id"`
	OriginAirportCode      string         `json:"origin_airport_code"`
	DestinationAirportCode string         `json:"destination_airport_code"`
	DepartureDateTime      time.Time      `json:"departure_date_time"`
	BookingClass           string         `json:"booking_class"`
	FlightNumber           string         `json:"flight_number"`
	FlightDuration         int32          `json:"flight_duration"`
	Enabled                bool           `json:"enabled"`
	SeatAvailable          int32          `json:"seat_available"`
	CurrencyCode           string         `json:"currency_code"`
	Price                  float64        `json:"price"`
	CreateBy               sql.NullString `json:"create_by"`
	CreateDate             sql.NullTime   `json:"create_date"`
	ModifyBy               sql.NullString `json:"modify_by"`
	ModifyDate             sql.NullTime   `json:"modify_date"`
}
