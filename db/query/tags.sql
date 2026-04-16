-- name: GetTagsByCafeID :many
select t."name",
    t.id
from cafes c
    inner join cafe_tags ct on ct.cafe_id = c.id
    inner join tags t on t.id = ct.tag_id
where c.id = $1;
