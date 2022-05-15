// Code generated by sqlc. DO NOT EDIT.
// source: category.sql

package db

import (
	"context"
)

const createCategory = `-- name: CreateCategory :one
INSERT INTO categories (
  name, featured_img
) VALUES (
  $1, $2
)
RETURNING id, name, featured_img, created_at, updated_at
`

type CreateCategoryParams struct {
	Name        string `json:"name"`
	FeaturedImg string `json:"featured_img"`
}

func (q *Queries) CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error) {
	row := q.db.QueryRowContext(ctx, createCategory, arg.Name, arg.FeaturedImg)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.FeaturedImg,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteCategory = `-- name: DeleteCategory :exec
DELETE FROM categories
WHERE id = $1
`

func (q *Queries) DeleteCategory(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCategory, id)
	return err
}

const getCategory = `-- name: GetCategory :one
SELECT id, name, featured_img, created_at, updated_at FROM categories
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCategory(ctx context.Context, id int64) (Category, error) {
	row := q.db.QueryRowContext(ctx, getCategory, id)
	var i Category
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.FeaturedImg,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTotalCategory = `-- name: GetTotalCategory :one
SELECT count(*) FROM categories
`

func (q *Queries) GetTotalCategory(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalCategory)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const listCategories = `-- name: ListCategories :many
SELECT id, name, featured_img, created_at, updated_at FROM categories
ORDER BY updated_at
LIMIT $1
OFFSET $2
`

type ListCategoriesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error) {
	rows, err := q.db.QueryContext(ctx, listCategories, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Category{}
	for rows.Next() {
		var i Category
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.FeaturedImg,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updateCategory = `-- name: UpdateCategory :exec
UPDATE categories
SET name = $2, featured_img = $3
WHERE id = $1
`

type UpdateCategoryParams struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	FeaturedImg string `json:"featured_img"`
}

func (q *Queries) UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error {
	_, err := q.db.ExecContext(ctx, updateCategory, arg.ID, arg.Name, arg.FeaturedImg)
	return err
}
