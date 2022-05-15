// Code generated by sqlc. DO NOT EDIT.
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
  username, bio, email, wallet_address, avatar,
  banner_img, ins_link, twitter_link, website_link
) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9
)
RETURNING id, username, bio, email, wallet_address, avatar, banner_img, ins_link, twitter_link, website_link, created_at, updated_at
`

type CreateUserParams struct {
	Username      string `json:"username"`
	Bio           string `json:"bio"`
	Email         string `json:"email"`
	WalletAddress string `json:"wallet_address"`
	Avatar        string `json:"avatar"`
	BannerImg     string `json:"banner_img"`
	InsLink       string `json:"ins_link"`
	TwitterLink   string `json:"twitter_link"`
	WebsiteLink   string `json:"website_link"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.Username,
		arg.Bio,
		arg.Email,
		arg.WalletAddress,
		arg.Avatar,
		arg.BannerImg,
		arg.InsLink,
		arg.TwitterLink,
		arg.WebsiteLink,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Bio,
		&i.Email,
		&i.WalletAddress,
		&i.Avatar,
		&i.BannerImg,
		&i.InsLink,
		&i.TwitterLink,
		&i.WebsiteLink,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE wallet_address = $1
`

func (q *Queries) DeleteUser(ctx context.Context, walletAddress string) error {
	_, err := q.db.ExecContext(ctx, deleteUser, walletAddress)
	return err
}

const getTotalUser = `-- name: GetTotalUser :one
SELECT count(*) FROM users
`

func (q *Queries) GetTotalUser(ctx context.Context) (int64, error) {
	row := q.db.QueryRowContext(ctx, getTotalUser)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, bio, email, wallet_address, avatar, banner_img, ins_link, twitter_link, website_link, created_at, updated_at FROM users
WHERE wallet_address = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, walletAddress string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, walletAddress)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Bio,
		&i.Email,
		&i.WalletAddress,
		&i.Avatar,
		&i.BannerImg,
		&i.InsLink,
		&i.TwitterLink,
		&i.WebsiteLink,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, bio, email, wallet_address, avatar, banner_img, ins_link, twitter_link, website_link, created_at, updated_at FROM users
ORDER BY updated_at
LIMIT $1
OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Bio,
			&i.Email,
			&i.WalletAddress,
			&i.Avatar,
			&i.BannerImg,
			&i.InsLink,
			&i.TwitterLink,
			&i.WebsiteLink,
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

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET bio = $2, website_link = $3, email = $4,
    avatar = $5, banner_img = $6, ins_link = $7,
    twitter_link = $8
WHERE wallet_address = $1
`

type UpdateUserParams struct {
	WalletAddress string `json:"wallet_address"`
	Bio           string `json:"bio"`
	WebsiteLink   string `json:"website_link"`
	Email         string `json:"email"`
	Avatar        string `json:"avatar"`
	BannerImg     string `json:"banner_img"`
	InsLink       string `json:"ins_link"`
	TwitterLink   string `json:"twitter_link"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser,
		arg.WalletAddress,
		arg.Bio,
		arg.WebsiteLink,
		arg.Email,
		arg.Avatar,
		arg.BannerImg,
		arg.InsLink,
		arg.TwitterLink,
	)
	return err
}
