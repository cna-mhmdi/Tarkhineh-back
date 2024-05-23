package api

import (
	db "github.com/cna-mhmdi/Tarkhineh-back/db/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	store  db.Store
	router *gin.Engine
}

func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.POST("/user", server.createUser)
	router.GET("/user/:username", server.getUser)
	//router.GET("/user", server.listUsers)
	router.DELETE("/user/:username", server.deleteUser)

	router.POST("/user/profile", server.createProfile)
	router.GET("/user/profile", server.getProfile)
	router.PUT("/user/profile", server.updateProfile)

	router.POST("/food", server.createFood)
	router.GET("/food/:name", server.getFood)
	router.GET("/food", server.listFoods)
	router.PUT("/food", server.updateFood)

	server.router = router
	return server
}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
