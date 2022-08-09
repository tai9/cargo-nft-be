package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/tai9/cargo-nft-be/constants"
)

type Store interface {
	Querier
	CreateNFTTx(ctx context.Context, arg CreateNFTParams) (CreateNFTTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database tranasction
func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}
	return tx.Commit()
}

type TransferTxParams struct {
	NftID           int64   `json:"nft_id"`
	FromUserID      int64   `json:"from_user_id"`
	ToUserID        int64   `json:"to_user_id"`
	Quantity        float64 `json:"quantity"`
	Event           string  `json:"event"`
	Token           string  `json:"token"`
	TransactionHash string  `json:"transaction_hash"`
}

type TransferTxResult struct {
	Transaction Transaction `json:"nft_id"`
	FromUser    User        `json:"from_user_id"`
	ToUser      User        `json:"to_user_id"`
}

// TransferTx performs a quantity transfer from one NFT to the other
// It creates a transfer record, add transaction and update NFT quantity within a single database transaction
// func (store *SQLStore) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
// 	var result TransferTxResult

// 	err := store.execTx(ctx, func(q *Queries) error {
// 		var err error

// 		result.Transaction, err = q.CreateTransaction(ctx, CreateTransactionParams{
// 			NftID:           arg.NftID,
// 			FromUserID:      arg.FromUserID,
// 			ToUserID:        arg.ToUserID,
// 			Quantity:        arg.Quantity,
// 			Event:           arg.Event,
// 			Token:           arg.Token,
// 			TransactionHash: arg.TransactionHash,
// 		})

// 		if err != nil {
// 			return err
// 		}

// 		if arg.FromUserID < arg.ToUserID {

// 		}

// 		return nil
// 	})

// 	return result, err
// }

// func addQuantity(
// 	ctx context.Context,
// 	q *Queries,
// 	userID1 int64,
// 	quantity1 float64,
// 	userID2 int64,
// 	quantity2 float64,
// ) (user1 User, user2 User, err error){

// }

type CreateNFTTxParams struct {
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

type CreateNFTTxResult struct {
	Nft         Nft         `json:"nft"`
	Transaction Transaction `json:"transaction"`
}

func (store *SQLStore) CreateNFTTx(ctx context.Context, arg CreateNFTParams) (CreateNFTTxResult, error) {
	var result CreateNFTTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.Nft, err = q.CreateNFT(ctx, CreateNFTParams{
			OwnerID:         arg.OwnerID,
			UserID:          arg.UserID,
			CollectionID:    arg.CollectionID,
			Name:            arg.Name,
			Description:     arg.Description,
			FeaturedImg:     arg.FeaturedImg,
			Supply:          arg.Supply,
			Views:           arg.Views,
			Favorites:       arg.Favorites,
			ContractAddress: arg.ContractAddress,
			TokenID:         arg.TokenID,
			TokenStandard:   arg.TokenStandard,
			Blockchain:      arg.Blockchain,
			Metadata:        arg.Metadata,
		})
		if err != nil {
			return err
		}

		result.Transaction, err = q.CreateTransaction(ctx, CreateTransactionParams{
			NftID:           result.Nft.ID,
			FromUserID:      result.Nft.UserID,
			ToUserID:        constants.NULL_ADDRESS_USER_ID,
			Quantity:        0,
			Event:           constants.MINTED,
			Token:           "",
			TransactionHash: "",
		})
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
