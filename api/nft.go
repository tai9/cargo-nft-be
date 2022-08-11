package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/tai9/cargo-nft-be/db/sqlc"
	"github.com/tai9/cargo-nft-be/token"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

type createNFTRequest struct {
	CollectionID  int64  `json:"collection_id" binding:"required"`
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description"`
	FeaturedImg   string `json:"featured_img" binding:"required"`
	Views         int64  `json:"views"`
	Favorites     string `json:"favorites"`
	TokenID       string `json:"token_id"`
	TokenStandard string `json:"token_standard"`
	Blockchain    string `json:"blockchain"`
	Metadata      string `json:"metadata"`
}

func (server *Server) createNFT(ctx *gin.Context) {
	var req createNFTRequest

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
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	nftCollection, err := server.thirdwebSdk.GetNFTCollection(collection.ContractAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	_, err = nftCollection.Mint(&thirdweb.NFTMetadataInput{
		Name:        req.Name,
		Description: req.Description,
		Image:       req.FeaturedImg,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	arg := db.CreateNFTParams{
		OwnerID:         user.ID,
		UserID:          user.ID,
		CollectionID:    req.CollectionID,
		Name:            req.Name,
		Description:     req.Description,
		FeaturedImg:     req.FeaturedImg,
		Supply:          1,
		Views:           req.Views,
		Favorites:       req.Favorites,
		ContractAddress: "",
		TokenID:         req.TokenID,
		TokenStandard:   req.TokenStandard,
		Blockchain:      req.Blockchain,
		Metadata:        req.Metadata,
	}

	nftTxResult, err := server.store.CreateNFTTx(ctx, arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, nftTxResult)
}

type updateNFTRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Supply      int64  `json:"supply"`
	FeaturedImg string `json:"featured_img"`
	Views       int64  `json:"views"`
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

type transferNFTRequest struct {
	ToAddress string `json:"to_address"`
}

func (server *Server) transferNFT(ctx *gin.Context) {
	var params NFTParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	var req transferNFTRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	user, err := server.store.GetUser(ctx, req.ToAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			user, err = server.store.CreateUser(ctx, db.CreateUserParams{
				WalletAddress: req.ToAddress,
			})
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errResponse(err))
				return
			}

		} else {
			ctx.JSON(http.StatusInternalServerError, errResponse(err))
			return
		}
	}

	nft, err := server.store.GetNFT(ctx, params.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	collection, err := server.store.GetCollection(ctx, nft.CollectionID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	nftCollection, err := server.thirdwebSdk.GetNFTCollection(collection.ContractAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	tokenId := 0
	_, err = nftCollection.Transfer(req.ToAddress, tokenId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	arg := db.UpdateNFTParams{
		ID:          params.ID,
		Name:        nft.Name,
		Description: nft.Description,
		FeaturedImg: nft.FeaturedImg,
		Supply:      nft.Supply,
		Views:       nft.Views,
		Favorites:   nft.Favorites,
		OwnerID:     user.ID,
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

	collection, err := server.store.GetCollection(ctx, 1)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	nftCollection, err := server.thirdwebSdk.GetNFTCollection(collection.ContractAddress)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	tokenId := 0
	_, err = nftCollection.Burn(tokenId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	err = server.store.DeleteNFT(ctx, params.ID)
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
	Page  int32            `json:"page"`
	Limit int32            `json:"limit"`
	Total int64            `json:"total"`
	Data  []db.ListNFTsRow `json:"data"`
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
	if req.Views == 0 {
		req.Views = nft.Views
	}
	if req.Favorites == "" {
		req.Favorites = nft.Favorites
	}
}
