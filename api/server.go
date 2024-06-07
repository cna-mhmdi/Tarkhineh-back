package api

import (
	"fmt"
	db "github.com/cna-mhmdi/Tarkhineh-back/db/sqlc"
	"github.com/cna-mhmdi/Tarkhineh-back/token"
	"github.com/cna-mhmdi/Tarkhineh-back/util"
	"github.com/cna-mhmdi/Tarkhineh-back/worker"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	router          *gin.Engine
	taskDistributor worker.TaskDistributor
}

func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:          config,
		store:           store,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {

	router := gin.Default()

	router.POST("/user", server.createUser)
	router.POST("/user/login", server.loginUser)
	router.POST("/token/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.GET("/user/:username", server.getUser)
	//router.GET("/user", server.listUsers)
	authRoutes.DELETE("/user/:username", server.deleteUser)

	authRoutes.POST("/user/profile", server.createProfile)
	authRoutes.GET("/user/profile", server.getProfile)
	authRoutes.PUT("/user/profile", server.updateProfile)

	router.POST("/food", server.createFood)
	authRoutes.GET("/food/:name", server.getFoodByName)
	authRoutes.GET("/food/getFoodById", server.getFoodById)
	authRoutes.GET("/food", server.listFoods)
	router.PUT("/food", server.updateFood)

	authRoutes.POST("/user/favorite", server.createFavoriteUser)
	authRoutes.GET("/user/favorite/:username", server.getFavoritesUser)
	authRoutes.DELETE("/user/favorite", server.deleteUserFavorite)

	authRoutes.POST("/user/address", server.createUserAddress)
	authRoutes.GET("/user/address/:username", server.getUserAddress)
	authRoutes.DELETE("/user/deleteAddress", server.deleteUserAddress)
	authRoutes.PUT("/user/address", server.updateUserAddress)

	server.router = router

}
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
