package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/piyush1146115/dummy-bank-golang/db/sqlc"
)

// Server serves HTTP requests for our banking service.
type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)
	router.GET("/accounts/:id", server.getAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
