package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/tai9/cargo-nft-be/db/sqlc"
	"github.com/tai9/cargo-nft-be/token"
	"github.com/thirdweb-dev/go-sdk/thirdweb"
)

type createCollectionRequest struct {
	Id             int64  `json:"id"`
	CategoryID     int64  `json:"category_id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	Description    string `json:"description"`
	Blockchain     string `json:"blockchain"`
	Owners         string `json:"owners"`
	PaymentToken   string `json:"payment_token" binding:"required"`
	CreatorEarning string `json:"creator_earning"`
	FeaturedImg    string `json:"featured_img" binding:"required"`
	BannerImg      string `json:"banner_img"`
	InsLink        string `json:"ins_link"`
	TwitterLink    string `json:"twitter_link"`
	WebsiteLink    string `json:"website_link"`
}

type updateCollectionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Owners      string `json:"owners"`
	FeaturedImg string `json:"featured_img"`
	BannerImg   string `json:"banner_img"`
	InsLink     string `json:"ins_link"`
	TwitterLink string `json:"twitter_link"`
	WebsiteLink string `json:"website_link"`
}

func (server *Server) createCollection(ctx *gin.Context) {
	var req createCollectionRequest

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

	contractAddress, err := server.thirdwebSdk.Deployer.DeployNFTCollection(&thirdweb.DeployNFTCollectionMetadata{
		Name:        req.Name,
		Description: req.Description,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	arg := db.CreateCollectionParams{
		UserID:          user.ID,
		Name:            req.Name,
		Description:     req.Description,
		Blockchain:      req.Blockchain,
		Owners:          req.Owners,
		PaymentToken:    req.PaymentToken,
		CreatorEarning:  req.CreatorEarning,
		FeaturedImg:     req.FeaturedImg,
		BannerImg:       req.BannerImg,
		InsLink:         req.InsLink,
		TwitterLink:     req.TwitterLink,
		WebsiteLink:     req.WebsiteLink,
		ContractAddress: contractAddress,
	}

	collection, err := server.store.CreateCollection(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	req.Id = collection.ID

	_, err = server.store.CreateCateCollection(ctx, db.CreateCateCollectionParams{
		CollectionID: collection.ID,
		CategoryID:   req.CategoryID,
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, req)
}

func (server *Server) updateCollection(ctx *gin.Context) {
	var params CollectionParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	var req updateCollectionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	collection, err := server.store.GetCollection(ctx, params.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	checkEmptyCollection(req, &collection)

	arg := db.UpdateCollectionParams{
		ID:          params.ID,
		Name:        req.Name,
		Description: req.Description,
		Owners:      req.Owners,
		FeaturedImg: req.FeaturedImg,
		BannerImg:   req.BannerImg,
		InsLink:     req.InsLink,
		TwitterLink: req.TwitterLink,
		WebsiteLink: req.WebsiteLink,
	}

	err = server.store.UpdateCollection(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

func (server *Server) deleteCollection(ctx *gin.Context) {
	var params CollectionParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	cateColections, err := server.store.ListCateCollections(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	for i := 0; i < len(cateColections)-1; i++ {
		if cateColections[i].CollectionID == params.ID {
			if err := server.store.DeleteCateCollection(ctx, cateColections[i].ID); err != nil {
				ctx.JSON(http.StatusInternalServerError, errResponse(err))
				return
			}

		}
	}

	err = server.store.DeleteCollection(ctx, params.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

type CollectionParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

type ListCollectionsParams struct {
	Page  int32 `form:"page" json:"page" binding:"min=1"`
	Limit int32 `form:"limit" json:"limit" binding:"min=0"`
}

type CollectionsResponse struct {
	db.GetCollectionRow
	TotalNFT int64 `json:"total_nft"`
}

type ListCollectionResponse struct {
	Page  int32                   `json:"page"`
	Limit int32                   `json:"limit"`
	Total int64                   `json:"total"`
	Data  []db.ListCollectionsRow `json:"data"`
}

func (server *Server) listCollection(ctx *gin.Context) {
	var req ListCollectionsParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListCollectionsParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}

	collections, err := server.store.ListCollections(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	totalCollections, err := server.store.GetTotalCollection(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	rsp := ListCollectionResponse{
		Page:  req.Page,
		Limit: req.Limit,
		Total: totalCollections,
		Data:  collections,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) getCollection(ctx *gin.Context) {
	var req CollectionParams

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	collection, err := server.store.GetCollection(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	totalNFT, err := server.store.GetTotalNFTByCollectionId(ctx, collection.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	rsp := CollectionsResponse{
		collection,
		totalNFT,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func checkEmptyCollection(req updateCollectionRequest, collection *db.GetCollectionRow) {
	if req.Name == "" {
		req.Name = collection.Name
	}
	if req.Description == "" {
		req.Description = collection.Description
	}
	if req.Owners == "" {
		req.Owners = collection.Owners
	}
	if req.FeaturedImg == "" {
		req.FeaturedImg = collection.FeaturedImg
	}
	if req.BannerImg == "" {
		req.BannerImg = collection.BannerImg
	}
	if req.InsLink == "" {
		req.InsLink = collection.InsLink
	}
	if req.TwitterLink == "" {
		req.TwitterLink = collection.TwitterLink
	}
	if req.WebsiteLink == "" {
		req.WebsiteLink = collection.WebsiteLink
	}
}
