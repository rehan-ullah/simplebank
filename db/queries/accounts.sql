-- name: CreateAccounts :one
INSERT INTO accounts (
  owner, balance,currency
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetAccounts :one
SELECT * FROM accounts
WHERE id = $1 LIMIT 1;

-- name: UpdateAccounts :one
UPDATE accounts
SET owner = $2,
balance = $3,
currency = $4
WHERE accounts.id = $1
RETURNING *;

-- name: GetAllAccounts :many
SELECT * FROM accounts
ORDER BY accounts.id LIMIT $1 OFFSET $2;

-- name: GetCountAllAccounts :one
SELECT COUNT(id) FROM accounts;

-- name: DeleteAccounts :exec
DELETE FROM accounts WHERE id = $1;