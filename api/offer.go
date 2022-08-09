package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tai9/cargo-nft-be/db/sqlc"
	"github.com/tai9/cargo-nft-be/token"
)

type createOfferRequest struct {
	NftID      int64     `json:"nft_id" binding:"required"`
	UsdPrice   float64   `json:"usd_price" binding:"required"`
	Token      string    `json:"token" binding:"required"`
	Quantity   float64   `json:"quantity" binding:"required"`
	Expiration time.Time `json:"expiration"`
}

func (server *Server) createOffer(ctx *gin.Context) {
	var req createOfferRequest

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

	arg := db.CreateOfferParams{
		UserID:          user.ID,
		NftID:           req.NftID,
		UsdPrice:        req.UsdPrice,
		Quantity:        req.Quantity,
		Token:           req.Token,
		Expiration:      req.Expiration,
		FloorDifference: 0, // TODO: calc with collection floor price

	}

	offer, err := server.store.CreateOffer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, offer)
}

type updateOfferRequest struct {
	UsdPrice        float64   `json:"usd_price"`
	Token           string    `json:"token"`
	Quantity        float64   `json:"quantity"`
	FloorDifference float64   `json:"floor_difference"`
	Expiration      time.Time `json:"expiration"`
}

func (server *Server) updateOffer(ctx *gin.Context) {
	var params OfferParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	var req updateOfferRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	offer, err := server.store.GetOffer(ctx, params.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	checkEmptyOffer(req, &offer)

	arg := db.UpdateOfferParams{
		ID:              params.ID,
		UsdPrice:        req.UsdPrice,
		Quantity:        req.Quantity,
		Expiration:      req.Expiration,
		FloorDifference: req.FloorDifference,
	}

	err = server.store.UpdateOffer(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, params)
}

func (server *Server) deleteOffer(ctx *gin.Context) {
	var params OfferParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err := server.store.DeleteOffer(ctx, params.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

type OfferParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type ListOffersParams struct {
	Page  int32 `form:"page" json:"page" binding:"min=1"`
	Limit int32 `form:"limit" json:"limit" binding:"min=0"`
}

type ListOfferResponse struct {
	Page  int32      `json:"page"`
	Limit int32      `json:"limit"`
	Total int64      `json:"total"`
	Data  []db.Offer `json:"data"`
}

func (server *Server) listOffer(ctx *gin.Context) {
	var req ListOffersParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListOffersParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}

	offers, err := server.store.ListOffers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	totalOffers, err := server.store.GetTotalOffer(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	rsp := ListOfferResponse{
		Page:  req.Page,
		Limit: req.Limit,
		Total: totalOffers,
		Data:  offers,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func checkEmptyOffer(req updateOfferRequest, offer *db.Offer) {
	if req.UsdPrice == 0 {
		req.UsdPrice = offer.UsdPrice
	}
	if req.Quantity == 0 {
		req.Quantity = offer.Quantity
	}
	if req.FloorDifference == 0 {
		req.FloorDifference = offer.FloorDifference
	}
	if req.Expiration == time.Now() {
		req.Expiration = offer.Expiration
	}
}
