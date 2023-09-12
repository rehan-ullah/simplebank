-- name: CreateEntries :one
INSERT INTO entries (
  account_id, amount
) VALUES (
  $1, $2
)
RETURNING *;

-- name: GetEntries :one
SELECT * FROM entries
WHERE id = $1 LIMIT 1;

-- name: UpdateEntries :one
UPDATE entries
SET account_id = $2,
amount = $3
WHERE id = $1
RETURNING *;

-- name: GetAllEntries :many
SELECT * FROM entries
ORDER BY entries.id;

-- name: DeleteEntries :exec
DELETE FROM entries WHERE id = $1;