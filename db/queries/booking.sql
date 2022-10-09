-- name: CreateBookingFlight :one
INSERT INTO booking_flights(
    booking_transaction_id, 
    sequence, 
    departure_airport_code, 
    arrival_airport_code, 
    departure_date_time, 
    arrival_date_time, 
    booking_class, 
    price, 
    currency_code
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: CreateBookingPassenger :one
INSERT INTO booking_passengers(
    booking_transaction_id, 
    title, 
    last_name, 
    first_name,
    ffp_number
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: CreateBookingTransaction :one
INSERT INTO booking_transactions(
    reservation_code, 
    status, 
    create_date
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: UpdateBookingTransactionStatus :one
UPDATE booking_transactions
SET 
    status=$2, 
    error_code=$3, 
    error_message=$4, 
    modify_date=$5
WHERE 
    id=$1
RETURNING *;

-- name: GetBookingTransactionsByReservationCode :many
SELECT * FROM booking_transactions
WHERE 
    reservation_code=$1
;