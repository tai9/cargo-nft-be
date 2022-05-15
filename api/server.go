package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tai9/cargo-nft-be/constants"
	db "github.com/tai9/cargo-nft-be/db/sqlc"
	"github.com/tai9/cargo-nft-be/token"
	"github.com/tai9/cargo-nft-be/utils"
)

// Server serves HTTP request for our services.
type Server struct {
	store      db.Store
	router     *gin.Engine
	config     utils.Config
	tokenMaker token.Maker
}

// NewServer creates a new HTTP server and setup routing.
func NewServer(config utils.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		store:      store,
		config:     config,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.Use(CORSMiddleware())

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "Hello world")
	})

	apiRoute := router.Group(constants.BASE_URL)

	apiRoute.POST("/connect-wallet", server.connectWallet)

	authRoutes := apiRoute.Group("/").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/profile", server.getUser)
	authRoutes.PATCH("/profile", server.updateUser)

	// remove on prod stage
	authRoutes.GET("/users", server.listUser)

	authRoutes.GET("/offers", server.listOffer)
	authRoutes.POST("/offers", server.createOffer)
	authRoutes.PATCH("/offers/:id", server.updateOffer)
	authRoutes.DELETE("/offers/:id", server.deleteOffer)

	authRoutes.GET("/collections", server.listCollection)
	authRoutes.POST("/collections", server.createCollection)
	authRoutes.PATCH("/collections/:id", server.updateCollection)
	authRoutes.DELETE("/collections/:id", server.deleteCollection)

	authRoutes.GET("/nfts", server.listNFT)
	authRoutes.POST("/nfts", server.createNFT)
	authRoutes.PATCH("/nfts/:id", server.updateNFT)
	authRoutes.DELETE("/nfts/:id", server.deleteNFT)

	authRoutes.GET("/transactions", server.listTransaction)
	authRoutes.POST("/transactions", server.createTransaction)
	authRoutes.DELETE("/transactions/:id", server.deleteTransaction)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
