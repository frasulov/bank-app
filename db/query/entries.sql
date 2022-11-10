-- name: GetEntry :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: GetEntries :many
SELECT * FROM entries
limit $1 offset $2;

-- name: CreateEntry :one
insert into entries (account_id, amount)
values  ($1,$2)
returning *;


-- name: DeleteEntry :exec
delete from entries
where id = $1;