-- name: GetUser :one
SELECT * FROM users
WHERE wallet_address = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetTotalUser :one
SELECT count(*) FROM users;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY updated_at
LIMIT $1
OFFSET $2;

-- name: CreateUser :one
INSERT INTO users (
  username, bio, email, wallet_address, avatar,
  banner_img, ins_link, twitter_link, website_link
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users
SET bio = $2, website_link = $3, email = $4,
    avatar = $5, banner_img = $6, ins_link = $7,
    twitter_link = $8
WHERE wallet_address = $1;

-- name: DeleteUser :exec
DELETE FROM users
WHERE wallet_address = $1;