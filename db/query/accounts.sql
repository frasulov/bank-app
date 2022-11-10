-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1 for no key update;

-- name: GetAccounts :many
SELECT * FROM accounts order by created_at desc
limit $1 offset $2;

-- name: CreateAccount :one
insert into accounts (owner, balance, currency)
values  ($1,$2,$3)
returning *;


-- name: UpdateAccount :one
update accounts set balance = $2 where id = $1 returning *;


-- name: DeleteAccount :exec
delete from accounts
where id = $1;