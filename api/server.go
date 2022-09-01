package api

import (
	db "simple-bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

// serves all http request for the app
type Server struct {
	store *db.Queries
	router *gin.Engine
}

// initialize server
func NewServer(store *db.Queries) *Server {
	server := &Server{store: store}
	router := gin.Default()
	server.router = router

	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.getUser)

	return server
}

// start a server on address
func (server *Server) Start(address string) error{
	return server.router.Run(address)
}

// custom error handler
func errorResponse(err error) gin.H {
	return gin.H{"An error occured": err.Error() }
}