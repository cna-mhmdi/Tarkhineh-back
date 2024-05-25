package api

import (
	"fmt"
	db "github.com/cna-mhmdi/Tarkhineh-back/db/sqlc"
	"github.com/cna-mhmdi/Tarkhineh-back/token"
	"github.com/cna-mhmdi/Tarkhineh-back/util"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {

	router := gin.Default()

	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)

	router.GET("/user/:username", server.getUser)
	//router.GET("/user", server.listUsers)
	router.DELETE("/user/:username", server.deleteUser)

	router.POST("/user/profile", server.createProfile)
	router.GET("/user/profile", server.getProfile)
	router.PUT("/user/profile", server.updateProfile)

	router.POST("/food", server.createFood)
	router.GET("/food/:name", server.getFood)
	router.GET("/food/getFoodById", server.getFoodById)
	router.GET("/food", server.listFoods)
	router.PUT("/food", server.updateFood)

	router.POST("/user/favorite", server.createFavoriteUser)
	router.GET("/user/favorite/:username", server.getFavoritesUser)
	router.DELETE("/user/favorite", server.deleteUserFavorite)

	router.POST("/user/address", server.createUserAddress)
	router.GET("/user/address/:username", server.getUserAddress)
	router.DELETE("/user/deleteAddress", server.deleteUserAddress)
	router.PUT("/user/address", server.updateUserAddress)

	server.router = router

}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
