-- name: GetTransaction :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: GetTotalTransaction :one
SELECT count(*) FROM transactions;

-- name: ListTransactions :many
SELECT * FROM transactions
ORDER BY updated_at
LIMIT $1
OFFSET $2;

-- name: CreateTransaction :one
INSERT INTO transactions (
  nft_id, event, token, quantity,
  from_user_id, to_user_id, transaction_hash
) VALUES (
  $1, $2, $3, $4, $5, $6, $7
)
RETURNING *;

-- name: DeleteTransaction :exec
DELETE FROM transactions
WHERE id = $1;