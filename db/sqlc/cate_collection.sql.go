// Code generated by sqlc. DO NOT EDIT.
// source: cate_collection.sql

package db

import (
	"context"
	"time"
)

const createCateCollection = `-- name: CreateCateCollection :one
INSERT INTO cate_collections (
  collection_id, category_id
) VALUES (
  $1, $2
)
RETURNING id, collection_id, category_id, created_at, updated_at
`

type CreateCateCollectionParams struct {
	CollectionID int64 `json:"collection_id"`
	CategoryID   int64 `json:"category_id"`
}

func (q *Queries) CreateCateCollection(ctx context.Context, arg CreateCateCollectionParams) (CateCollection, error) {
	row := q.db.QueryRowContext(ctx, createCateCollection, arg.CollectionID, arg.CategoryID)
	var i CateCollection
	err := row.Scan(
		&i.ID,
		&i.CollectionID,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCateCollection = `-- name: DeleteCateCollection :exec
DELETE FROM cate_collections
WHERE id = $1
`

func (q *Queries) DeleteCateCollection(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCateCollection, id)
	return err
}

const getCateCollection = `-- name: GetCateCollection :one
SELECT id, collection_id, category_id, created_at, updated_at FROM cate_collections
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCateCollection(ctx context.Context, id int64) (CateCollection, error) {
	row := q.db.QueryRowContext(ctx, getCateCollection, id)
	var i CateCollection
	err := row.Scan(
		&i.ID,
		&i.CollectionID,
		&i.CategoryID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTotalCateCollection = `-- name: GetTotalCateCollection :one
SELECT count(*) FROM cate_collections
`

func (q *Queries) GetTotalCateCollection(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalCateCollection)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const listCateCollections = `-- name: ListCateCollections :many
SELECT cate_collections.id, cate_collections.collection_id, cate_collections.category_id, cate_collections.created_at, cate_collections.updated_at, categories."name" category_name FROM cate_collections, categories, collections
WHERE cate_collections.category_id = categories.id
AND cate_collections.collection_id = collections.id
`

type ListCateCollectionsRow struct {
	ID           int64     `json:"id"`
	CollectionID int64     `json:"collection_id"`
	CategoryID   int64     `json:"category_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	CategoryName string    `json:"category_name"`
}

func (q *Queries) ListCateCollections(ctx context.Context) ([]ListCateCollectionsRow, error) {
	rows, err := q.db.QueryContext(ctx, listCateCollections)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListCateCollectionsRow{}
	for rows.Next() {
		var i ListCateCollectionsRow
		if err := rows.Scan(
			&i.ID,
			&i.CollectionID,
			&i.CategoryID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CategoryName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCateCollection = `-- name: UpdateCateCollection :exec
UPDATE cate_collections
SET collection_id = $2, category_id = $3
WHERE id = $1
`

type UpdateCateCollectionParams struct {
	ID           int64 `json:"id"`
	CollectionID int64 `json:"collection_id"`
	CategoryID   int64 `json:"category_id"`
}

func (q *Queries) UpdateCateCollection(ctx context.Context, arg UpdateCateCollectionParams) error {
	_, err := q.db.ExecContext(ctx, updateCateCollection, arg.ID, arg.CollectionID, arg.CategoryID)
	return err
}