package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/tai9/cargo-nft-be/db/sqlc"
)

type ListCategoryParams struct {
	Page  int32 `form:"page" json:"page" binding:"min=1" default:"1"`
	Limit int32 `form:"limit" json:"limit" binding:"min=0" default:"10"`
}

type ListCategoryResponse struct {
	Page  int32                       `json:"page"`
	Limit int32                       `json:"limit"`
	Total int64                       `json:"total"`
	Data  []db.ListCateCollectionsRow `json:"data"`
}

func (server *Server) listCategory(ctx *gin.Context) {
	categories, err := server.store.ListCateCollections(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, categories)

}
