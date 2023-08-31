package api

import (
	"github.com/gin-gonic/gin"

	db "github.com/xlal1234/simplebank/db/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store *db.Store

	router *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(store *db.Store) (*Server, error) {

	server := &Server{store: store}

	// server.setupRouter()
	router := gin.Default()
	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccounts)

	server.router = router
	return server, nil
}

// func (server *Server) setupRouter() {
// 	router := gin.Default()

// 	router.POST("/users", server.createUser)
// 	router.POST("/users/login", server.loginUser)
// 	router.POST("/tokens/renew_access", server.renewAccessToken)

// 	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
// 	authRoutes.POST("/accounts", server.createAccount)
// 	authRoutes.GET("/accounts/:id", server.getAccount)
// 	authRoutes.GET("/accounts", server.listAccounts)

// 	authRoutes.POST("/transfers", server.createTransfer)

// 	server.router = router
// }

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
