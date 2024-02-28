-- @sql postgresql
-- name: GetUserById :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: GetUserByName :one
SELECT *
FROM users
WHERE name = $1
LIMIT 1;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;

-- name: ListUsers :many
select id, name, email, role, hashed_password, created_at, password_changed_at, img_file
from users;

-- name: CreateUser :one
INSERT INTO users (name, email, role, hashed_password, created_at, password_changed_at)
VALUES ($1, $2, $3, $4, now(), '0001-01-01 00:00:00Z'::timestamp)
RETURNING *;

-- name: DeleteUser :exec
DELETE
FROM users
WHERE id = $1;

-- name: AttachePhoto :exec
update users
set img_file=$1
where id = $2;

-- name: GetImagePath :one
select COALESCE(img_file, '')
from users
where id = $1;

-- name: UpdateUser :exec
update users
set name=$1,
    email=$2,
    role=$3
where id = $4;

-- name: UpdateUserPassword :exec
update users
set hashed_password=$1
where id = $2;