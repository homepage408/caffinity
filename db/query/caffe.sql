-- name: GetCafeByID :one
SELECT *
FROM cafes
WHERE id = $1;

-- name: CreateCafe :one
INSERT INTO cafes (name, city, rating)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateCafe :one
UPDATE cafes
SET name = $2,
    city = $3,
    rating = $4
WHERE id = $1
RETURNING *;

-- name: GetCafesByCity :many
SELECT *
FROM cafes
WHERE city ILIKE $1
LIMIT $2 OFFSET $3;

-- name: GetAllCafes :many
SELECT *
FROM cafes
LIMIT $1 OFFSET $2;