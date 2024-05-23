package api

import (
	"database/sql"
	db "github.com/cna-mhmdi/Tarkhineh-back/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createFoodRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int32  `json:"price" binding:"required,min=1000"`
	Rate        int32  `json:"rate" binding:"required,min=5,max=10"`
	Discount    int32  `json:"discount" binding:"required,min=-1,max=100"`
	FoodTag     string `json:"food_tag" binding:"required"`
}

func (server *Server) createFood(ctx *gin.Context) {
	var req createFoodRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateFoodParams{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Rate:        req.Rate,
		Discount:    req.Discount,
		FoodTag:     req.FoodTag,
	}

	food, err := server.store.CreateFood(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, food)
}

type getFoodRequest struct {
	Name string `uri:"name" binding:"required"`
}

func (server *Server) getFood(ctx *gin.Context) {
	var req getFoodRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	food, err := server.store.GetFood(ctx, req.Name)
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

type listFoodRequest struct {
	PageId   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listFoods(ctx *gin.Context) {
	var req listFoodRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListFoodsParams{
		Limit:  req.PageSize,
		Offset: (req.PageId - 1) * req.PageSize,
	}

	foods, err := server.store.ListFoods(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, foods)
}

type updateFoodRequest struct {
	ID          int64  `json:"id" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description" binding:"required"`
	Price       int32  `json:"price" binding:"required,min=1000"`
	Rate        int32  `json:"rate" binding:"required,min=1,max=5"`
	Discount    int32  `json:"discount" binding:"required,min=0,max=100"`
	FoodTag     string `json:"food_tag" binding:"required"`
}

func (server *Server) updateFood(ctx *gin.Context) {
	var req updateFoodRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateFoodParams{
		ID:          req.ID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Rate:        req.Rate,
		Discount:    req.Discount,
		FoodTag:     req.FoodTag,
	}

	profile, err := server.store.UpdateFood(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profile)
}
