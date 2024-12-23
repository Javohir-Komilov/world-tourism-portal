-- name: CreateDriver :one
INSERT INTO drivers (user_id, vehicle, availability_status)
VALUES (?, ?, ?) RETURNING id;

-- name: UpdateDriver :exec
UPDATE drivers
SET vehicle = ?, availability_status = ?
WHERE id = ?;

-- name: GetDriverById :one
SELECT * FROM drivers WHERE id = ? LIMIT 1;

-- name: ListDrivers :many
SELECT * FROM drivers ORDER BY id;

-- name: CreateHotel :one
INSERT INTO hotels (manager_id, name, location, rating)
VALUES (?, ?, ?, ?) RETURNING id;

-- name: UpdateHotel :exec
UPDATE hotels
SET name = ?, location = ?, rating = ?
WHERE id = ?;

-- name: GetHotelById :one
SELECT * FROM hotels WHERE id = ? LIMIT 1;

-- name: ListHotels :many
SELECT * FROM hotels ORDER BY name;

-- name: CreateTourPackage :one
INSERT INTO tour_packages (agent_id, name, description, price, start_date, end_date)
VALUES (?, ?, ?, ?, ?, ?) RETURNING id;

-- name: UpdateTourPackage :exec
UPDATE tour_packages
SET name = ?, description = ?, price = ?, start_date = ?, end_date = ?
WHERE id = ?;

-- name: GetTourPackageById :one
SELECT * FROM tour_packages WHERE id = ? LIMIT 1;

-- name: ListTourPackages :many
SELECT * FROM tour_packages ORDER BY start_date;

-- name: CreateUser :one
INSERT INTO users (username, password_hash, role)
VALUES (?, ?, ?) RETURNING id;

-- name: UpdateUser :exec
UPDATE users
SET username = ?, password_hash = ?, role = ?
WHERE id = ?;

-- name: GetUserById :one
SELECT * FROM users WHERE id = ? LIMIT 1;

-- name: GetUserByUsername :one
SELECT * FROM users WHERE username = ? LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users ORDER BY id;

-- name: CreateHotelBooking :one
INSERT INTO hotel_bookings (hotel_id, tourist_id, check_in, check_out)
VALUES (?, ?, ?, ?) RETURNING id;

-- name: UpdateHotelBooking :exec
UPDATE hotel_bookings
SET check_in = ?, check_out = ?
WHERE id = ?;

-- name: GetHotelBookingById :one
SELECT * FROM hotel_bookings WHERE id = ? LIMIT 1;

-- name: ListHotelBookingsByTourist :many
SELECT * FROM hotel_bookings WHERE tourist_id = ? ORDER BY check_in;

-- name: CreateTripOrder :one
INSERT INTO trip_orders (driver_id, tourist_id, destination, status)
VALUES (?, ?, ?, ?) RETURNING id;

-- name: UpdateTripOrderStatus :exec
UPDATE trip_orders
SET status = ?
WHERE id = ?;

-- name: GetTripOrderById :one
SELECT * FROM trip_orders WHERE id = ? LIMIT 1;

-- name: ListTripOrdersByDriver :many
SELECT * FROM trip_orders WHERE driver_id = ? ORDER BY created_at DESC;

-- name: ListTripOrdersByTourist :many
SELECT * FROM trip_orders WHERE tourist_id = ? ORDER BY created_at DESC;
