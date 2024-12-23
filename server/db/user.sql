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
