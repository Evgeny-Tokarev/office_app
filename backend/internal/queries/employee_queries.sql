-- @sql postgresql
-- name: GetEmployee :one
SELECT *
FROM employees
WHERE id = $1
LIMIT 1;

-- name: ListEmployees :many
select id, name, age, office_id, created_at, updated_at, img_file
from employees
where office_id = $1;

-- name: CreateEmployee :one
insert into employees (name, age, created_at, updated_at, office_id)
values ($1, $2, NOW(), NOW(), $3)
RETURNING *;

-- name: DeleteEmployee :exec
DELETE
FROM employees
WHERE id = $1;

-- name: AttachePhoto :exec
update employees
set img_file=$1
where id = $2;

-- name: GetImagePath :one
select COALESCE(img_file, '')
from employees
where id = $1;

-- name: UpdateEmployee :exec
update employees
set name=$1,
    age=$2,
    updated_at=NOW()
where id = $3;





