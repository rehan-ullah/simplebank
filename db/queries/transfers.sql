-- name: CreateTransfers :one
INSERT INTO transfers (
  from_account_id, to_account_id, amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransfers :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: UpdateTransfers :one
UPDATE transfers
SET from_account_id = $2,
to_account_id = $3,
amount = $4
WHERE id = $1
RETURNING *;

-- name: GetAllTransfers :many
SELECT * FROM transfers
ORDER BY transfers.id;

-- name: DeleteTansfers :exec
DELETE FROM transfers WHERE id = $1;