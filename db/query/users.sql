-- name: GetUser :one
SELECT *
FROM users
WHERE username = $1 LIMIT 1;

-- name: CreateUser :one
insert into users (username, password, full_name, email)
values ($1, $2, $3, $4) returning *;
