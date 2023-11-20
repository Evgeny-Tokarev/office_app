-- @sql postgresql
-- name: GetOffice :one
SELECT * FROM offices
WHERE id = $1 LIMIT 1;

-- name: ListOffices :many
SELECT * FROM offices
ORDER BY id;

-- name: CreateOffice :one
insert into offices (name, address, created_at, updated_at)
values ($1, $2, now(), now())
RETURNING *;

-- name: DeleteOffice :exec
DELETE FROM offices
WHERE id = $1;

-- name: AttachePhoto :exec
update offices set img_file=$1
where id=$2;

-- name: GetImagePath :one
select COALESCE(img_file, '') from offices where id = $1;

-- name: UpdateOffice :exec
update offices set name=$1, address=$2, updated_at=$3
where id = $4;