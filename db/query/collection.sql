-- name: GetCollection :one
SELECT collections.*, users.username created_by FROM collections, users
WHERE collections.user_id = users.id
AND collections.id = $1
LIMIT 1;

-- name: GetTotalCollection :one
SELECT count(*) FROM collections;

-- name: ListCollections :many
SELECT collections.*, users.username created_by FROM collections, users
WHERE collections.user_id = users.id
ORDER BY updated_at
LIMIT $1
OFFSET $2;

-- name: CreateCollection :one
INSERT INTO collections (
  user_id, name, description, blockchain,
  owners, payment_token, creator_earning, featured_img,
  banner_img, ins_link, twitter_link, website_link
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
)
RETURNING *;

-- name: UpdateCollection :exec
UPDATE collections
SET name = $2, description = $3, owners = $4, featured_img = $5,
    banner_img = $6, ins_link = $7, twitter_link = $8, website_link = $9
WHERE id = $1;

-- name: DeleteCollection :exec
DELETE FROM collections
WHERE id = $1;