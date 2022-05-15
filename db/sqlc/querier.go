// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateCateCollection(ctx context.Context, arg CreateCateCollectionParams) (CateCollection, error)
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateCollection(ctx context.Context, arg CreateCollectionParams) (Collection, error)
	CreateCryptoCurrency(ctx context.Context, arg CreateCryptoCurrencyParams) (CryptoCurrency, error)
	CreateNFT(ctx context.Context, arg CreateNFTParams) (Nft, error)
	CreateOffer(ctx context.Context, arg CreateOfferParams) (Offer, error)
	CreateTransaction(ctx context.Context, arg CreateTransactionParams) (Transaction, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCateCollection(ctx context.Context, id int64) error
	DeleteCategory(ctx context.Context, id int64) error
	DeleteCollection(ctx context.Context, id int64) error
	DeleteCryptoCurrency(ctx context.Context, code string) error
	DeleteNFT(ctx context.Context, id int64) error
	DeleteOffer(ctx context.Context, id int64) error
	DeleteTransaction(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, walletAddress string) error
	GetCateCollection(ctx context.Context, id int64) (CateCollection, error)
	GetCategory(ctx context.Context, id int64) (Category, error)
	GetCollection(ctx context.Context, id int64) (Collection, error)
	GetCryptoCurrency(ctx context.Context, code string) (CryptoCurrency, error)
	GetNFT(ctx context.Context, id int64) (Nft, error)
	GetOffer(ctx context.Context, id int64) (Offer, error)
	GetTotalCateCollection(ctx context.Context) (int64, error)
	GetTotalCategory(ctx context.Context) (int64, error)
	GetTotalCollection(ctx context.Context) (int64, error)
	GetTotalCryptoCurrency(ctx context.Context) (int64, error)
	GetTotalNFT(ctx context.Context) (int64, error)
	GetTotalOffer(ctx context.Context) (int64, error)
	GetTotalTransaction(ctx context.Context) (int64, error)
	GetTotalUser(ctx context.Context) (int64, error)
	GetTransaction(ctx context.Context, id int64) (Transaction, error)
	GetUser(ctx context.Context, walletAddress string) (User, error)
	ListCateCollections(ctx context.Context) ([]ListCateCollectionsRow, error)
	ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error)
	ListCollections(ctx context.Context, arg ListCollectionsParams) ([]Collection, error)
	ListCryptoCurrencies(ctx context.Context, arg ListCryptoCurrenciesParams) ([]CryptoCurrency, error)
	ListNFTs(ctx context.Context, arg ListNFTsParams) ([]Nft, error)
	ListOffers(ctx context.Context, arg ListOffersParams) ([]Offer, error)
	ListTransactions(ctx context.Context, arg ListTransactionsParams) ([]Transaction, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	UpdateCateCollection(ctx context.Context, arg UpdateCateCollectionParams) error
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) error
	UpdateCollection(ctx context.Context, arg UpdateCollectionParams) error
	UpdateCryptoCurrency(ctx context.Context, arg UpdateCryptoCurrencyParams) error
	UpdateNFT(ctx context.Context, arg UpdateNFTParams) error
	UpdateOffer(ctx context.Context, arg UpdateOfferParams) error
	UpdateUser(ctx context.Context, arg UpdateUserParams) error
}

var _ Querier = (*Queries)(nil)
