-- name: GetSession :one
SELECT *
FROM sessions
WHERE id = $1 LIMIT 1;

-- name: CreateSession :one
insert into sessions (id, username, refresh_token, user_agent, client_ip, is_blocked, expires_at)
values ($1, $2, $3, $4, $5, $6, $7) returning *;
