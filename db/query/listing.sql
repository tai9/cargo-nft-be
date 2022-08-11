-- name: GetListing :one
SELECT * FROM listings
WHERE id = $1 LIMIT 1;

-- name: GetTotalListing :one
SELECT count(*) FROM listings;

-- name: ListListings :many
SELECT * FROM listings
ORDER BY updated_at
LIMIT $1
OFFSET $2;

-- name: CreateListing :one
INSERT INTO listings (
  user_id, nft_id, usd_unit_price, quantity, usd_price, expiration, token, from_user_id, listing_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateListing :exec
UPDATE listings
SET usd_price = $2, quantity = $3, usd_price = $4, expiration = $5, from_user_id = $6
WHERE id = $1;

-- name: DeleteListing :exec
DELETE FROM listings
WHERE id = $1;