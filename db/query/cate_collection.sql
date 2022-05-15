-- name: GetCateCollection :one
SELECT * FROM cate_collections
WHERE id = $1 LIMIT 1;

-- name: GetTotalCateCollection :one
SELECT count(*) FROM cate_collections;

-- name: ListCateCollections :many
SELECT cate_collections.*, categories."name" category_name FROM cate_collections, categories, collections
WHERE cate_collections.category_id = categories.id
AND cate_collections.collection_id = collections.id;

-- name: CreateCateCollection :one
INSERT INTO cate_collections (
  collection_id, category_id
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdateCateCollection :exec
UPDATE cate_collections
SET collection_id = $2, category_id = $3
WHERE id = $1;

-- name: DeleteCateCollection :exec
DELETE FROM cate_collections
WHERE id = $1;