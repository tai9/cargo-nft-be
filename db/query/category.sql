-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1 LIMIT 1;

-- name: GetTotalCategory :one
SELECT count(*) FROM categories;

-- name: ListCategories :many
SELECT * FROM categories
ORDER BY updated_at
LIMIT $1
OFFSET $2;

-- name: CreateCategory :one
INSERT INTO categories (
  name, featured_img
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateCategory :exec
UPDATE categories
SET name = $2, featured_img = $3
WHERE id = $1;

-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1;