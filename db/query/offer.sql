-- name: GetOffer :one
SELECT * FROM offers
WHERE id = $1 LIMIT 1;

-- name: GetTotalOffer :one
SELECT count(*) FROM offers;

-- name: ListOffers :many
SELECT * FROM offers
ORDER BY updated_at
LIMIT $1
OFFSET $2;

-- name: CreateOffer :one
INSERT INTO offers (
  user_id, nft_id, usd_price, quantity, floor_difference, expiration
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING *;

-- name: UpdateOffer :exec
UPDATE offers
SET usd_price = $2, quantity = $3, floor_difference = $4, expiration = $5
WHERE id = $1;

-- name: DeleteOffer :exec
DELETE FROM offers
WHERE id = $1;