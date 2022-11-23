-- name: GetAccount :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1 for no key update;

-- name: GetAccounts :many
SELECT * FROM accounts
         where owner = $1
         order by created_at desc
limit $2 offset $3;

-- name: CreateAccount :one
insert into accounts (owner, balance, currency)
values  ($1,$2,$3)
returning *;


-- name: UpdateAccount :one
update accounts set balance = $2 where id = $1 returning *;

-- name: AddAccountBalance :one
update accounts set balance = balance + $2 where id = $1 returning *;

-- name: DeleteAccount :exec
delete from accounts
where id = $1;

-- name: GetFullAccountInfo :one
select jsonb_build_object('id', a.id, 'name', a.owner, 'balance', a.balance, 'currency', a.currency, 'transactions', coalesce(x.transactions, '[]'))
from accounts a
         left join lateral (
    select jsonb_agg(jsonb_build_object('amount', t.amount, 'created_at', t.created_at)) as transactions
    from transfers t where t.from_account_id = a.id
        ) as x on true where a.id= $1;
