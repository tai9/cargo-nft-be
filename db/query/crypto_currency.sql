-- name: GetCryptoCurrency :one
SELECT * FROM crypto_currencies
WHERE code = $1 LIMIT 1;

-- name: GetTotalCryptoCurrency :one
SELECT count(*) FROM crypto_currencies;

-- name: ListCryptoCurrencies :many
SELECT * FROM crypto_currencies
ORDER BY updated_at
LIMIT $1
OFFSET $2;

-- name: CreateCryptoCurrency :one
INSERT INTO crypto_currencies (
  name, code, price
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: UpdateCryptoCurrency :exec
UPDATE crypto_currencies
SET name = $2, code = $3, price = $4
WHERE code = $1;

-- name: DeleteCryptoCurrency :exec
DELETE FROM crypto_currencies
WHERE code = $1;