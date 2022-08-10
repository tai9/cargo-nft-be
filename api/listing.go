package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tai9/cargo-nft-be/db/sqlc"
	"github.com/tai9/cargo-nft-be/token"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

type createListingRequest struct {
	NftID        int64     `json:"nft_id" binding:"required"`
	CollectionID int64     `json:"collection_id" binding:"required"`
	UsdPrice     float64   `json:"usd_price" binding:"required"`
	Token        string    `json:"token" binding:"required"`
	Expiration   time.Time `json:"expiration"`
}

func (server *Server) createListing(ctx *gin.Context) {
	var req createListingRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUser(ctx, authPayload.WalletAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	collection, err := server.store.GetCollection(ctx, req.CollectionID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	marketplace, err := server.thirdwebSdk.GetMarketplace(server.config.MarketplaceContractAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	listing := &thirdweb.NewDirectListing{
		AssetContractAddress:     collection.ContractAddress,
		TokenId:                  0,     //TODO: change to actual nft token id
		ListingDurationInSeconds: 86400, //TODO: using expiredTime from request payload
		Quantity:                 1,
		BuyoutPricePerToken:      req.UsdPrice,
		CurrencyContractAddress:  authPayload.WalletAddress,
	}
	listingId, err := marketplace.CreateListing(listing)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	fmt.Println("listingId: ", listingId)

	arg := db.CreateListingParams{
		UserID:     user.ID,
		NftID:      req.NftID,
		UsdPrice:   req.UsdPrice,
		Quantity:   1,
		Token:      req.Token,
		Expiration: req.Expiration,
		FromUserID: user.ID,
	}

	listingCreated, err := server.store.CreateListing(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, listingCreated)
}

type updateListingRequest struct {
	UsdPrice   float64   `json:"usd_price"`
	Token      string    `json:"token"`
	Quantity   float64   `json:"quantity"`
	Expiration time.Time `json:"expiration"`
}

func (server *Server) updateListing(ctx *gin.Context) {
	var params ListingParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	var req updateListingRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	listing, err := server.store.GetListing(ctx, params.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	checkEmptyListing(req, &listing)

	arg := db.UpdateListingParams{
		ID:         params.ID,
		UsdPrice:   req.UsdPrice,
		Quantity:   req.Quantity,
		Expiration: req.Expiration,
	}

	err = server.store.UpdateListing(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, params)
}

func (server *Server) deleteListing(ctx *gin.Context) {
	var params ListingParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err := server.store.DeleteListing(ctx, params.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

type ListingParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type ListListingsParams struct {
	Page  int32 `form:"page" json:"page" binding:"min=1"`
	Limit int32 `form:"limit" json:"limit" binding:"min=0"`
}

type ListListingResponse struct {
	Page  int32        `json:"page"`
	Limit int32        `json:"limit"`
	Total int64        `json:"total"`
	Data  []db.Listing `json:"data"`
}

func (server *Server) listListing(ctx *gin.Context) {
	var req ListListingsParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	// arg := db.ListListingsParams{
	// 	Limit:  req.Limit,
	// 	Offset: (req.Page - 1) * req.Limit,
	// }

	marketplace, err := server.thirdwebSdk.GetMarketplace(server.config.MarketplaceContractAddress)
	if err != nil {
		fmt.Println("hiiiii")
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	listings, err := marketplace.GetActiveListings(&thirdweb.MarketplaceFilter{
		TokenContract: "0x80e2a899BA8970e0c9b153412CA3F7646DED0285",
	})
	if err != nil {
		fmt.Println("hahah")
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	// listings, err := server.store.ListListings(ctx, arg)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, errResponse(err))
	// 	return
	// }

	// totalListings, err := server.store.GetTotalListing(ctx)
	// if err != nil {
	// 	ctx.JSON(http.StatusInternalServerError, errResponse(err))
	// 	return
	// }

	// rsp := ListListingResponse{
	// 	Page:  req.Page,
	// 	Limit: req.Limit,
	// 	Total: totalListings,
	// 	Data:  listings,
	// }

	ctx.JSON(http.StatusOK, listings)
}

func checkEmptyListing(req updateListingRequest, listing *db.Listing) {
	if req.UsdPrice == 0 {
		req.UsdPrice = listing.UsdPrice
	}
	if req.Quantity == 0 {
		req.Quantity = listing.Quantity
	}
	if req.Expiration == time.Now() {
		req.Expiration = listing.Expiration
	}
}
