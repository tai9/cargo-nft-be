// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type CateCollection struct {
	ID           int64     `json:"id"`
	CollectionID int64     `json:"collection_id"`
	CategoryID   int64     `json:"category_id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Category struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	FeaturedImg string    `json:"featured_img"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Collection struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"user_id"`
	Name           string    `json:"name"`
	Description    string    `json:"description"`
	Blockchain     string    `json:"blockchain"`
	Owners         string    `json:"owners"`
	PaymentToken   string    `json:"payment_token"`
	CreatorEarning string    `json:"creator_earning"`
	FeaturedImg    string    `json:"featured_img"`
	BannerImg      string    `json:"banner_img"`
	InsLink        string    `json:"ins_link"`
	TwitterLink    string    `json:"twitter_link"`
	WebsiteLink    string    `json:"website_link"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type CryptoCurrency struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Nft struct {
	ID              int64     `json:"id"`
	UserID          int64     `json:"user_id"`
	CollectionID    int64     `json:"collection_id"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	FeaturedImg     string    `json:"featured_img"`
	Supply          int32     `json:"supply"`
	Views           string    `json:"views"`
	Favorites       string    `json:"favorites"`
	ContractAddress string    `json:"contract_address"`
	TokenID         string    `json:"token_id"`
	TokenStandard   string    `json:"token_standard"`
	Blockchain      string    `json:"blockchain"`
	Metadata        string    `json:"metadata"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Offer struct {
	ID              int64     `json:"id"`
	UserID          int64     `json:"user_id"`
	NftID           int64     `json:"nft_id"`
	Quantity        float64   `json:"quantity"`
	UsdPrice        float64   `json:"usd_price"`
	Token           string    `json:"token"`
	FloorDifference string    `json:"floor_difference"`
	Expiration      time.Time `json:"expiration"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Transaction struct {
	ID              int64     `json:"id"`
	NftID           int64     `json:"nft_id"`
	Event           string    `json:"event"`
	Quantity        float64   `json:"quantity"`
	Token           string    `json:"token"`
	FromUserID      int64     `json:"from_user_id"`
	ToUserID        int64     `json:"to_user_id"`
	TransactionHash string    `json:"transaction_hash"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type User struct {
	ID            int64     `json:"id"`
	Username      string    `json:"username"`
	Bio           string    `json:"bio"`
	Email         string    `json:"email"`
	WalletAddress string    `json:"wallet_address"`
	Avatar        string    `json:"avatar"`
	BannerImg     string    `json:"banner_img"`
	InsLink       string    `json:"ins_link"`
	TwitterLink   string    `json:"twitter_link"`
	WebsiteLink   string    `json:"website_link"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
