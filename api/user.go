package api

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	db "github.com/tai9/cargo-nft-be/db/sqlc"
	"github.com/tai9/cargo-nft-be/token"
)

type connectWalletRequest struct {
	WalletAddress string `json:"wallet_address" binding:"required,alphanum"`
}

type connectWalletResponse struct {
	UserID        int64     `json:"user_id"`
	Username      string    `json:"username"`
	WalletAddress string    `json:"wallet_address"`
	AccessToken   string    `json:"access_token"`
	ExpiredAt     time.Time `json:"expired_at"`
}

func (server *Server) connectWallet(ctx *gin.Context) {
	var req connectWalletRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
	}

	user := db.User{}
	var err error

	user, err = server.store.GetUser(ctx, req.WalletAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			// create user with the wallet address was connected
			arg := db.CreateUserParams{}
			arg.WalletAddress = req.WalletAddress
			user, err = server.store.CreateUser(ctx, arg)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, errResponse(err))
				return
			}
		} else {
			ctx.JSON(http.StatusUnauthorized, errResponse(err))
			return
		}
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(req.WalletAddress, server.config.AccessTokenDuration)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	rsp := &connectWalletResponse{
		UserID:        user.ID,
		Username:      user.Username,
		WalletAddress: user.WalletAddress,
		AccessToken:   accessToken,
		ExpiredAt:     accessPayload.ExpiredAt,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func (server *Server) getUser(ctx *gin.Context) {
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	user, err := server.store.GetUser(ctx, authPayload.WalletAddress)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, user)
}

type updateUserRequest struct {
	Bio         string `json:"bio"`
	WebsiteLink string `json:"website_link"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
	BannerImg   string `json:"banner_img"`
	InsLink     string `json:"ins_link"`
	TwitterLink string `json:"twitter_link"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	var err error

	if err = ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	user, err := server.store.GetUser(ctx, authPayload.WalletAddress)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	checkEmptyUser(req, &user)

	arg := db.UpdateUserParams{
		WalletAddress: authPayload.WalletAddress,
		Bio:           req.Bio,
		WebsiteLink:   req.WebsiteLink,
		Email:         req.Email,
		Avatar:        req.Avatar,
		BannerImg:     req.BannerImg,
		InsLink:       req.InsLink,
		TwitterLink:   req.TwitterLink,
	}

	err = server.store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	ctx.Status(http.StatusOK)
}

type ListUsersParams struct {
	Page  int32 `form:"page" json:"page" binding:"min=1"`
	Limit int32 `form:"limit" json:"limit" binding:"min=0"`
}

type ListUserResponse struct {
	Page  int32     `json:"page"`
	Limit int32     `json:"limit"`
	Total int64     `json:"total"`
	Data  []db.User `json:"data"`
}

func (server *Server) listUser(ctx *gin.Context) {
	var req ListUsersParams
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errResponse(err))
		return
	}

	arg := db.ListUsersParams{
		Limit:  req.Limit,
		Offset: (req.Page - 1) * req.Limit,
	}
	users, err := server.store.ListUsers(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	totalUser, err := server.store.GetTotalUser(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errResponse(err))
		return
	}

	// TODO: update user response
	rsp := ListUserResponse{
		Page:  req.Page,
		Limit: req.Limit,
		Total: totalUser,
		Data:  users,
	}

	ctx.JSON(http.StatusOK, rsp)
}

func checkEmptyUser(req updateUserRequest, user *db.User) {
	if req.Bio == "" {
		req.Bio = user.Bio
	}
	if req.WebsiteLink == "" {
		req.WebsiteLink = user.WebsiteLink
	}
	if req.Email == "" {
		req.Email = user.Email
	}
	if req.Avatar == "" {
		req.Avatar = user.Avatar
	}
	if req.BannerImg == "" {
		req.BannerImg = user.BannerImg
	}
	if req.InsLink == "" {
		req.InsLink = user.InsLink
	}
	if req.TwitterLink == "" {
		req.TwitterLink = user.TwitterLink
	}
}
