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
