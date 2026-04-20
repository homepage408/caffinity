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

-- name: GetFacilitiesByCafeIDs :many
select 
  c.id as cafe_id,
  f.id,
  f.name
from cafes c
inner join cafe_facilities cf on cf.cafe_id = c.id
inner join facilities f on f.id = cf.facility_id 
where c.id = any($1::int[]);

-- name: GetTagsByCafeIDs :many
select 
  c.id as cafe_id,
  t.id as tag_id,
  t.name
from cafes c
inner join cafe_tags ct on ct.cafe_id = c.id
inner join tags t on t.id = ct.tag_id
where c.id = any($1::int[]);

-- name: GetMenusByCafeIDs :many
select
  m.id,
  m.cafe_id,
  m.name,
  m.price,
  m.strength,
  m.is_safe,
  m.description,
  m.image
from menus m
inner join cafes c on c.id = m.cafe_id
where m.cafe_id = any($1::int[]);