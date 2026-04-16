-- name: GetFacilitiesByCafeID :many
select f.id,
    f.name
from cafes c
    inner join cafe_facilities cf on cf.cafe_id = c.id
    inner join facilities f on f.id = cf.facility_id
where c.id = $1;