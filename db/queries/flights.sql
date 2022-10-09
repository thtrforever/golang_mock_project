-- name: CreateFlight :one
INSERT INTO flights(
	origin_airport_code, 
    destination_airport_code, 
    departure_date_time, 
    booking_class,
    flight_number,
    flight_duration,
    enabled, 
    seat_available, 
    currency_code, 
    price, 
    create_by, 
    create_date
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING *;

-- name: GetFlightById :one
SELECT * FROM flights
WHERE 
    id=$1
;

-- name: GetFlightsByFields :many
SELECT * FROM flights
WHERE 
    origin_airport_code=$3
    AND destination_airport_code=$4
    AND date_trunc('day', departure_date_time)<=$5
    AND (date_trunc('day', departure_date_time) + interval '1 day' - interval '1 second')>=$5
    AND flight_number=$6
    AND booking_class=$7
    AND seat_available>0
ORDER BY 
    origin_airport_code ASC,
    destination_airport_code ASC,
    departure_date_time ASC
LIMIT $1
OFFSET $2
;

-- name: GetAvailableFlightsByFields :many
SELECT * FROM flights
WHERE 
    enabled='1'
    AND seat_available>0
    AND origin_airport_code=$3
    AND destination_airport_code=$4
    AND date_trunc('day', departure_date_time)<=$5
    AND (date_trunc('day', departure_date_time) + interval '1 day' - interval '1 second')>=$5
    AND (flight_number=$6 OR $6='')
    AND (booking_class=ANY(sqlc.narg('BookingClasses')::varchar[]) OR array_length(sqlc.narg('BookingClasses')::varchar[], 1) IS NULL)
ORDER BY 
    origin_airport_code ASC,
    destination_airport_code ASC,
    flight_number ASC,
    departure_date_time ASC
LIMIT $1
OFFSET $2
;

-- name: UpdateFlight :one
UPDATE flights
SET 
    origin_airport_code=$2,
    destination_airport_code=$3,
    departure_date_time=$4, 
    booking_class=$5, 
    flight_number=$6,
    flight_duration=$7, 
    enabled=$8, 
    seat_available=$9, 
    currency_code=$10, 
    price=$11, 
    modify_by=$12, 
    modify_date=$13
WHERE 
    id=$1
RETURNING *;

-- name: UpdateFlightSeatAvailable :one
UPDATE flights
SET 
    seat_available=$2, 
    modify_by=$3, 
    modify_date=$4
WHERE 
    id=$1
RETURNING *;

-- name: DeleteFlight :exec
DELETE FROM flights
WHERE
    id=$1
;