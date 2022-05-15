-- name: GetNFT :one
SELECT * FROM nfts
WHERE id = $1 LIMIT 1;

-- name: GetTotalNFT :one
SELECT count(*) FROM nfts
WHERE LOWER(nfts."name") LIKE sqlc.arg(search)::varchar;

-- name: GetTotalNFTByCollectionId :one
SELECT COUNT(nfts.collection_id) FROM collections LEFT JOIN nfts
ON nfts.collection_id = collections.id
GROUP BY collections.id HAVING collections.id = $1;

-- name: ListNFTs :many
SELECT nfts.*, collections."name" collection_name FROM nfts, collections
WHERE nfts.collection_id = collections.id AND LOWER(nfts."name") LIKE sqlc.arg(search)::varchar
ORDER BY updated_at
LIMIT $1
OFFSET $2;

-- name: CreateNFT :one
INSERT INTO nfts (
  user_id, collection_id, name, description, featured_img, supply,
  views, favorites, contract_address, token_id, token_standard,
  blockchain, metadata
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING *;

-- name: UpdateNFT :exec
UPDATE nfts
SET name = $2, description = $3, supply = $4, featured_img = $5,
    views = $6, favorites = $7
WHERE id = $1;

-- name: DeleteNFT :exec
DELETE FROM nfts
WHERE id = $1;