// Code generated by sqlc. DO NOT EDIT.
// source: nft.sql

package db

import (
	"context"
)

const createNFT = `-- name: CreateNFT :one
INSERT INTO nfts (
  user_id, collection_id, name, description, featured_img, supply,
  views, favorites, contract_address, token_id, token_standard,
  blockchain, metadata
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13
)
RETURNING id, user_id, collection_id, name, description, featured_img, supply, views, favorites, contract_address, token_id, token_standard, blockchain, metadata, created_at, updated_at
`

type CreateNFTParams struct {
	UserID          int64  `json:"user_id"`
	CollectionID    int64  `json:"collection_id"`
	Name            string `json:"name"`
	Description     string `json:"description"`
	FeaturedImg     string `json:"featured_img"`
	Supply          int32  `json:"supply"`
	Views           string `json:"views"`
	Favorites       string `json:"favorites"`
	ContractAddress string `json:"contract_address"`
	TokenID         string `json:"token_id"`
	TokenStandard   string `json:"token_standard"`
	Blockchain      string `json:"blockchain"`
	Metadata        string `json:"metadata"`
}

func (q *Queries) CreateNFT(ctx context.Context, arg CreateNFTParams) (Nft, error) {
	row := q.db.QueryRowContext(ctx, createNFT,
		arg.UserID,
		arg.CollectionID,
		arg.Name,
		arg.Description,
		arg.FeaturedImg,
		arg.Supply,
		arg.Views,
		arg.Favorites,
		arg.ContractAddress,
		arg.TokenID,
		arg.TokenStandard,
		arg.Blockchain,
		arg.Metadata,
	)
	var i Nft
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CollectionID,
		&i.Name,
		&i.Description,
		&i.FeaturedImg,
		&i.Supply,
		&i.Views,
		&i.Favorites,
		&i.ContractAddress,
		&i.TokenID,
		&i.TokenStandard,
		&i.Blockchain,
		&i.Metadata,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteNFT = `-- name: DeleteNFT :exec
DELETE FROM nfts
WHERE id = $1
`

func (q *Queries) DeleteNFT(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteNFT, id)
	return err
}

const getNFT = `-- name: GetNFT :one
SELECT id, user_id, collection_id, name, description, featured_img, supply, views, favorites, contract_address, token_id, token_standard, blockchain, metadata, created_at, updated_at FROM nfts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetNFT(ctx context.Context, id int64) (Nft, error) {
	row := q.db.QueryRowContext(ctx, getNFT, id)
	var i Nft
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.CollectionID,
		&i.Name,
		&i.Description,
		&i.FeaturedImg,
		&i.Supply,
		&i.Views,
		&i.Favorites,
		&i.ContractAddress,
		&i.TokenID,
		&i.TokenStandard,
		&i.Blockchain,
		&i.Metadata,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTotalNFT = `-- name: GetTotalNFT :one
SELECT count(*) FROM nfts
WHERE LOWER(nfts."name") LIKE $1::varchar
`

func (q *Queries) GetTotalNFT(ctx context.Context, search string) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalNFT, search)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getTotalNFTByCollectionId = `-- name: GetTotalNFTByCollectionId :one
SELECT COUNT(nfts.collection_id) FROM collections LEFT JOIN nfts
ON nfts.collection_id = collections.id
GROUP BY collections.id HAVING collections.id = $1
`

func (q *Queries) GetTotalNFTByCollectionId(ctx context.Context, id int64) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalNFTByCollectionId, id)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const listNFTs = `-- name: ListNFTs :many
SELECT id, user_id, collection_id, name, description, featured_img, supply, views, favorites, contract_address, token_id, token_standard, blockchain, metadata, created_at, updated_at FROM nfts
WHERE LOWER(nfts."name") LIKE $3::varchar
ORDER BY updated_at
LIMIT $1
OFFSET $2
`

type ListNFTsParams struct {
	Limit  int32  `json:"limit"`
	Offset int32  `json:"offset"`
	Search string `json:"search"`
}

func (q *Queries) ListNFTs(ctx context.Context, arg ListNFTsParams) ([]Nft, error) {
	rows, err := q.db.QueryContext(ctx, listNFTs, arg.Limit, arg.Offset, arg.Search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Nft{}
	for rows.Next() {
		var i Nft
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.CollectionID,
			&i.Name,
			&i.Description,
			&i.FeaturedImg,
			&i.Supply,
			&i.Views,
			&i.Favorites,
			&i.ContractAddress,
			&i.TokenID,
			&i.TokenStandard,
			&i.Blockchain,
			&i.Metadata,
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

const updateNFT = `-- name: UpdateNFT :exec
UPDATE nfts
SET name = $2, description = $3, supply = $4, featured_img = $5,
    views = $6, favorites = $7
WHERE id = $1
`

type UpdateNFTParams struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Supply      int32  `json:"supply"`
	FeaturedImg string `json:"featured_img"`
	Views       string `json:"views"`
	Favorites   string `json:"favorites"`
}

func (q *Queries) UpdateNFT(ctx context.Context, arg UpdateNFTParams) error {
	_, err := q.db.ExecContext(ctx, updateNFT,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Supply,
		arg.FeaturedImg,
		arg.Views,
		arg.Favorites,
	)
	return err
}