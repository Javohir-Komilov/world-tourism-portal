-- name: CreateDriver :one
INSERT INTO drivers (user_id, vehicle, availability_status)
VALUES (:user_id, :vehicle, :availability_status)
RETURNING id;

-- name: UpdateDriver :one
UPDATE drivers
SET vehicle = :vehicle,
    availability_status = :availability_status
WHERE id = :id
RETURNING id;

-- name: CreateHotel :one
INSERT INTO hotels (manager_id, name, location, rating)
VALUES (:manager_id, :name, :location, :rating)
RETURNING id;

-- name: UpdateHotel :one
UPDATE hotels
SET name = :name,
    location = :location,
    rating = :rating
WHERE id = :id
RETURNING id;

-- name: CreateTourPackage :one
INSERT INTO tour_packages (agent_id, name, description, price, start_date, end_date)
VALUES (:agent_id, :name, :description, :price, :start_date, :end_date)
RETURNING id;

-- name: UpdateTourPackage :one
UPDATE tour_packages
SET name = :name,
    description = :description,
    price = :price,
    start_date = :start_date,
    end_date = :end_date
WHERE id = :id
RETURNING id;

-- name: CreateUser :one
INSERT INTO users (username, password_hash, role)
VALUES (:username, :password_hash, :role)
RETURNING id;

-- name: UpdateUser :one
UPDATE users
SET username = :username,
    password_hash = :password_hash,
    role = :role
WHERE id = :id
RETURNING id;
