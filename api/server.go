package api

import (
	db "github.com/Sahas001/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/accounts", server.createAccount)

	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	router.DELETE("/accounts/:id", server.deleteAccount)

	router.PUT("/accounts/:id", server.updateAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func okayResponse(message string) gin.H {
	return gin.H{"message": message}
}
