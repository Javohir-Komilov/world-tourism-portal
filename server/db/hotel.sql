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
