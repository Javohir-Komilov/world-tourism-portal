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
