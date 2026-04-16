-- name: GetMenusByCafeID :many
select m.id,
    m.cafe_id,
    m.name,
    m.price,
    m.strength,
    m.is_safe,
    m.description,
    m.image
from menus m
where m.cafe_id = $1;