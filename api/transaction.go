package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/tai9/cargo-nft-be/db/sqlc"
	"github.com/tai9/cargo-nft-be/token"
)

type createTransactionRequest struct {
	NftID    int64   `json:"nft_id" binding:"required"`
	Event    string  `json:"event" binding:"required"`
	Token    string  `json:"token"`
	Quantity float64 `json:"quantity"`
	ToUserID int64   `json:"to_user_id" binding:"required"`
}

func (server *Server) createTransaction(ctx *gin.Context) {
	var req createTransactionRequest

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

	arg := db.CreateTransactionParams{
		NftID:           req.NftID,
		Event:           req.Event,
		Token:           req.Token,
		Quantity:        req.Quantity,
		FromUserID:      user.ID,
		ToUserID:        req.ToUserID,
		TransactionHash: "",
	}

	transaction, err := server.store.CreateTransaction(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, transaction)
}

type TransactionParams struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) deleteTransaction(ctx *gin.Context) {
	var params TransactionParams
	if err := ctx.ShouldBindUri(&params); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	err := server.store.DeleteTransaction(ctx, params.ID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

type ListTransactionsParams struct {
	Page  int32 `form:"page" json:"page" binding:"min=1"`
	Limit int32 `form:"limit" json:"limit" binding:"min=0"`
}

type ListTransactionResponse struct {
	Page  int32            `json:"page"`
	Limit int32            `json:"limit"`
	Total int64            `json:"total"`
	Data  []db.Transaction `json:"data"`
}

func (server *Server) listTransaction(ctx *gin.Context) {
	var req ListTransactionsParams

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListTransactionsParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}

	transactions, err := server.store.ListTransactions(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	totalTransactions, err := server.store.GetTotalTransaction(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	rsp := ListTransactionResponse{
		Page:  req.Page,
		Limit: req.Limit,
		Total: totalTransactions,
		Data:  transactions,
	}

	ctx.JSON(http.StatusOK, rsp)
}
