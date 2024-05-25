package api

import (
	"database/sql"
	db "github.com/cna-mhmdi/Tarkhineh-back/db/sqlc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type createUserAddressRequest struct {
	Username    string `json:"username" binding:"required"`
	AddressLine string `json:"address_line" binding:"required"`
	AddressTag  string `json:"address_tag" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func (server *Server) createUserAddress(ctx *gin.Context) {
	var req createUserAddressRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateAddressParams{
		Username:    req.Username,
		AddressLine: req.AddressLine,
		AddressTag:  req.AddressTag,
		PhoneNumber: req.PhoneNumber,
	}

	address, err := server.store.CreateAddress(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, address)
}

type getUserAddressRequest struct {
	Username string `uri:"username" binding:"required"`
}

func (server *Server) getUserAddress(ctx *gin.Context) {
	var req getUserAddressRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	address, err := server.store.GetAddresses(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, address)
}

type deleteUserAddressRequest struct {
	Username string `json:"username" binding:"required"`
	ID       int64  `json:"id" binding:"required,min=0"`
}

func (server *Server) deleteUserAddress(ctx *gin.Context) {
	var req deleteUserAddressRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.DeleteAddressParams{
		Username: req.Username,
		ID:       req.ID,
	}

	result, err := server.store.DeleteAddress(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if rowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "no matching record found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"response": "user's address is successfully deleted"})
}

type updateUserAddressRequest struct {
	ID          int64  `json:"id" binding:"required" binding:"required"`
	AddressLine string `json:"address_line" binding:"required"`
	AddressTag  string `json:"address_tag" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"required"`
}

func (server *Server) updateUserAddress(ctx *gin.Context) {
	var req updateUserAddressRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateAddressParams{
		ID:          req.ID,
		AddressLine: req.AddressLine,
		AddressTag:  req.AddressTag,
		PhoneNumber: req.PhoneNumber,
	}

	address, err := server.store.UpdateAddress(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, address)
}
