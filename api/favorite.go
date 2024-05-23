package api

import (
	"database/sql"
	"errors"
	db "github.com/cna-mhmdi/Tarkhineh-back/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createFavoriteRequest struct {
	Username string `json:"username" binding:"required"`
	FoodId   int64  `json:"food_id" binding:"required"`
}

func (server *Server) createFavoriteUser(ctx *gin.Context) {
	var req createFavoriteRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFavoriteParams{
		Username: req.Username,
		FoodID:   req.FoodId,
	}

	food, err := server.store.CreateFavorite(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, food)
}

type getFavoritesUserRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) getFavoritesUser(ctx *gin.Context) {
	var req getFavoritesUserRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	food, err := server.store.GetFavorites(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, food)
}

type deleteUserFavoriteRequest struct {
	Username string `json:"username" binding:"required"`
	FoodId   int64  `json:"food_id" binding:"required,min=0"`
}

func (server *Server) deleteUserFavorite(ctx *gin.Context) {
	var req deleteUserFavoriteRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteFavoriteParams{
		Username: req.Username,
		FoodID:   req.FoodId,
	}
	err := server.store.DeleteFavorite(ctx, arg)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, "user's favorite is successfully deleted")
}
