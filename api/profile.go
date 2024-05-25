package api

import (
	"database/sql"
	"errors"
	db "github.com/cna-mhmdi/Tarkhineh-back/db/sqlc"
	"github.com/cna-mhmdi/Tarkhineh-back/token"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"net/http"
)

type createProfileRequest struct {
	Username    string `json:"username" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"` //e164 format should be added for phone number format validation
	BirthDay    string `json:"birthday" binding:"required"`
	NickName    string `json:"nickname" binding:"required"`
}

func (server *Server) createProfile(ctx *gin.Context) {
	var req createProfileRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	arg := db.CreateProfileParams{
		Username:    authPayload.Username,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Birthday:    req.BirthDay,
		Nickname:    req.NickName,
	}

	Profile, err := server.store.CreateProfile(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, Profile)
}

type getUserProfileRequest struct {
	Username string `json:"username" binding:"required,alpha"`
}

func (server *Server) getProfile(ctx *gin.Context) {
	var req getUserProfileRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	profile, err := server.store.GetProfile(ctx, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if profile.Username != authPayload.Username {
		err := errors.New("profile doesn't belong to the authenticated user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profile)
}

// update api's should be updated and username should be added for authorization
type updateUserProfileRequest struct {
	ID          int64  `json:"id" binding:"required"`
	FirstName   string `json:"first_name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Email       string `json:"email" binding:"required,email"`
	PhoneNumber string `json:"phone_number" binding:"required"` //e164 format should be added for phone number format validation
	BirthDay    string `json:"birthday" binding:"required"`
	NickName    string `json:"nickname" binding:"required"`
}

func (server *Server) updateProfile(ctx *gin.Context) {
	var req updateUserProfileRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateProfileParams{
		ID:          req.ID,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		Birthday:    req.BirthDay,
		Nickname:    req.NickName,
	}

	profile, err := server.store.UpdateProfile(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, profile)
}
