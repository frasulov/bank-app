-- name: GetTransfer :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: GetTransfers :many
SELECT * FROM transfers
limit $1 offset $2;

-- name: CreateTransfer :one
insert into transfers (from_account_id, to_account_id, amount)
values  ($1,$2,$3)
returning *;


-- name: DeleteTransfer :exec
delete from transfers
where id = $1;