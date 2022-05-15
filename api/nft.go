package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/tai9/cargo-nft-be/db/sqlc"
)

type createNFTRequest struct {
	UserID          int64  `json:"user_id" binding:"required"`
	CollectionID    int64  `json:"collection_id" binding:"required"`
	Name            string `json:"name" binding:"required"`
	Description     string `json:"description"`
	FeaturedImg     string `json:"featured_img" binding:"required"`
	Supply          int32  `json:"supply" binding:"required"`
	Views           string `json:"views"`
	Favorites       string `json:"favorites"`
	ContractAddress string `json:"contract_address" binding:"required"`
	TokenID         string `json:"token_id"`
	TokenStandard   string `json:"token_standard"`
	Blockchain      string `json:"blockchain"`
	Metadata        string `json:"metadata"`
}

func (server *Server) createNFT(ctx *gin.Context) {
	var req createNFTRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.CreateNFTParams{
		UserID:          req.UserID,
		CollectionID:    req.CollectionID,
		Name:            req.Name,
		Description:     req.Description,
		FeaturedImg:     req.FeaturedImg,
		Supply:          req.Supply,
		Views:           req.Views,
		Favorites:       req.Favorites,
		ContractAddress: req.ContractAddress,
		TokenID:         req.TokenID,
		TokenStandard:   req.TokenStandard,
		Blockchain:      req.Blockchain,
		Metadata:        req.Metadata,
	}

	NFT, err := server.store.CreateNFT(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, NFT)
}

type updateNFTRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Supply      int32  `json:"supply"`
	FeaturedImg string `json:"featured_img"`
	Views       string `json:"views"`
	Favorites   string `json:"favorites"`
}

func (server *Server) updateNFT(ctx *gin.Context) {
	var params NFTParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	var req updateNFTRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	nft, err := server.store.GetNFT(ctx, params.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	checkEmptyNFT(req, &nft)

	arg := db.UpdateNFTParams{
		ID:          params.ID,
		Name:        req.Name,
		Description: req.Description,
		FeaturedImg: req.FeaturedImg,
		Supply:      req.Supply,
		Views:       req.Views,
		Favorites:   req.Favorites,
	}

	err = server.store.UpdateNFT(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, params)
}

func (server *Server) deleteNFT(ctx *gin.Context) {
	var params NFTParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err := server.store.DeleteNFT(ctx, params.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

type NFTParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type ListNFTsParams struct {
	Page   int32  `form:"page" json:"page" binding:"min=1" default:"1"`
	Limit  int32  `form:"limit" json:"limit" binding:"min=0"`
	Search string `form:"search" json:"search" default:""`
}

type ListNFTResponse struct {
	Page  int32    `json:"page"`
	Limit int32    `json:"limit"`
	Total int64    `json:"total"`
	Data  []db.Nft `json:"data"`
}

func (server *Server) listNFT(ctx *gin.Context) {
	var req ListNFTsParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	searchQuery := fmt.Sprintf("%%%s%%", req.Search)

	arg := db.ListNFTsParams{
		Search: searchQuery,
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}

	nfts, err := server.store.ListNFTs(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	totalNFTs, err := server.store.GetTotalNFT(ctx, searchQuery)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	rsp := ListNFTResponse{
		Page:  req.Page,
		Limit: req.Limit,
		Total: totalNFTs,
		Data:  nfts,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func checkEmptyNFT(req updateNFTRequest, nft *db.Nft) {
	if req.Name == "0" {
		req.Name = nft.Name
	}
	if req.Description == "0" {
		req.Description = nft.Description
	}
	if req.Supply == 0 {
		req.Supply = nft.Supply
	}
	if req.FeaturedImg == "" {
		req.FeaturedImg = nft.FeaturedImg
	}
	if req.Views == "" {
		req.Views = nft.Views
	}
	if req.Favorites == "" {
		req.Favorites = nft.Favorites
	}

}
