// Code generated by sqlc. DO NOT EDIT.
// source: listing.sql

package db

import (
	"context"
	"time"
)

const createListing = `-- name: CreateListing :one
INSERT INTO listings (
  user_id, nft_id, usd_unit_price, quantity, usd_price, expiration, token, from_user_id, listing_id
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, user_id, nft_id, from_user_id, quantity, usd_price, usd_unit_price, token, expiration, created_at, updated_at, listing_id
`

type CreateListingParams struct {
	UserID       int64     `json:"user_id"`
	NftID        int64     `json:"nft_id"`
	UsdUnitPrice float64   `json:"usd_unit_price"`
	Quantity     float64   `json:"quantity"`
	UsdPrice     float64   `json:"usd_price"`
	Expiration   time.Time `json:"expiration"`
	Token        string    `json:"token"`
	FromUserID   int64     `json:"from_user_id"`
	ListingID    int32     `json:"listing_id"`
}

func (q *Queries) CreateListing(ctx context.Context, arg CreateListingParams) (Listing, error) {
	row := q.db.QueryRowContext(ctx, createListing,
		arg.UserID,
		arg.NftID,
		arg.UsdUnitPrice,
		arg.Quantity,
		arg.UsdPrice,
		arg.Expiration,
		arg.Token,
		arg.FromUserID,
		arg.ListingID,
	)
	var i Listing
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.NftID,
		&i.FromUserID,
		&i.Quantity,
		&i.UsdPrice,
		&i.UsdUnitPrice,
		&i.Token,
		&i.Expiration,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ListingID,
	)
	return i, err
}

const deleteListing = `-- name: DeleteListing :exec
DELETE FROM listings
WHERE id = $1
`

func (q *Queries) DeleteListing(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteListing, id)
	return err
}

const getListing = `-- name: GetListing :one
SELECT id, user_id, nft_id, from_user_id, quantity, usd_price, usd_unit_price, token, expiration, created_at, updated_at, listing_id FROM listings
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetListing(ctx context.Context, id int64) (Listing, error) {
	row := q.db.QueryRowContext(ctx, getListing, id)
	var i Listing
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.NftID,
		&i.FromUserID,
		&i.Quantity,
		&i.UsdPrice,
		&i.UsdUnitPrice,
		&i.Token,
		&i.Expiration,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ListingID,
	)
	return i, err
}

const getTotalListing = `-- name: GetTotalListing :one
SELECT count(*) FROM listings
`

func (q *Queries) GetTotalListing(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalListing)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const listListings = `-- name: ListListings :many
SELECT id, user_id, nft_id, from_user_id, quantity, usd_price, usd_unit_price, token, expiration, created_at, updated_at, listing_id FROM listings
ORDER BY updated_at
LIMIT $1
OFFSET $2
`

type ListListingsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListListings(ctx context.Context, arg ListListingsParams) ([]Listing, error) {
	rows, err := q.db.QueryContext(ctx, listListings, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Listing{}
	for rows.Next() {
		var i Listing
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.NftID,
			&i.FromUserID,
			&i.Quantity,
			&i.UsdPrice,
			&i.UsdUnitPrice,
			&i.Token,
			&i.Expiration,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ListingID,
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

const updateListing = `-- name: UpdateListing :exec
UPDATE listings
SET usd_price = $2, quantity = $3, usd_price = $4, expiration = $5, from_user_id = $6
WHERE id = $1
`

type UpdateListingParams struct {
	ID         int64     `json:"id"`
	UsdPrice   float64   `json:"usd_price"`
	Quantity   float64   `json:"quantity"`
	UsdPrice_2 float64   `json:"usd_price_2"`
	Expiration time.Time `json:"expiration"`
	FromUserID int64     `json:"from_user_id"`
}

func (q *Queries) UpdateListing(ctx context.Context, arg UpdateListingParams) error {
	_, err := q.db.ExecContext(ctx, updateListing,
		arg.ID,
		arg.UsdPrice,
		arg.Quantity,
		arg.UsdPrice_2,
		arg.Expiration,
		arg.FromUserID,
	)
	return err
}
