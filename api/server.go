package api

import (
	"github.com/gin-gonic/gin"
	db "swiftiesoft.com/simplebank/db/sqlc"
)

// Server serves HTTP requests for the project
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// NewServer create a new HTTP server & setup routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	routers := gin.Default()

	routers.POST("/accounts", server.CreateAccount)
	routers.GET("/accounts/:id", server.GetAccount)
	routers.GET("/getAllAccounts", server.GetAllAccounts)

	server.router = routers
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
